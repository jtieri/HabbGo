package transport

import (
	"bufio"
	"bytes"
	"context"
	"log/slog"
	"net"
	"strings"
	"sync"

	"github.com/jtieri/habbgo/protocol"
	"github.com/jtieri/habbgo/protocol/incoming"
	"github.com/jtieri/habbgo/protocol/outgoing"
	"github.com/jtieri/habbgo/services"
)

// Session represents a user's underlying network connection to the server.
// It is used to write and read data to and from the wire.
type Session struct {
	log              *slog.Logger
	id               string
	address          string
	connection       net.Conn
	buff             *buffer
	disconnectedChan chan DisconnectedSession
	outgoingRegistry *outgoing.OutgoingRegistry
	incomingRegistry *incoming.IncomingRegistry
}

// buffer is the buffered Writer used to write data to a Session's connection.
type buffer struct {
	mux  sync.Mutex
	buff *bufio.Writer
}

// Write writes a slice of bytes to the underling buffered writer.
// It returns the number of bytes written to the writer, as well as an error if the entire slice was not written.
func (b *buffer) Write(p []byte) (int, error) {
	b.mux.Lock()
	defer b.mux.Unlock()
	return b.buff.Write(p)
}

// Flush flushes the contents of the underlying buffered writer over the network.
func (b *buffer) Flush() error {
	b.mux.Lock()
	defer b.mux.Unlock()
	return b.buff.Flush()
}

func NewSession(
	log *slog.Logger,
	connection net.Conn,
	address string,
	id string,
	disconnectedChan chan DisconnectedSession,
) *Session {
	return &Session{
		log:              log,
		connection:       connection,
		buff:             &buffer{mux: sync.Mutex{}, buff: bufio.NewWriter(connection)},
		id:               id,
		address:          address,
		disconnectedChan: disconnectedChan,
		outgoingRegistry: outgoing.NewOutgoingRegistry(),
		incomingRegistry: incoming.NewIncomingRegistry(),
	}
}

// listen instantiates a connection to the game client,
// then reads and validates incoming data from the underlying network connection.
// It builds incoming packet objects and finds the appropriate handler for valid packets.
func (s *Session) listen(ctx context.Context, queues *services.ServiceQueues) {
	address := s.address

	reader := bufio.NewReader(s.connection)

	// Send packet with Base64 header @@ to initialize connection with client.
	s.Send(outgoing.Hello())

	// Listen for incoming packets from a player's session.
	for {
		// Attempt to read three bytes,
		// client->server packets in FUSEv0.2.0 begin with 3 byte Base64 encoded packet length.
		encodedLen, err := readPacketLength(reader)
		if err != nil {
			// If the network connection is closed, it's because the server closed the Session
			// which means we don't need to log again or call session.Close
			if strings.Contains(err.Error(), "use of closed network connection") {
				return
			}

			s.log.Error(
				"Error reading encoded packet length from session",
				"session_address", address,
				"error", err,
			)

			s.DisconnectSession()
			return
		}

		// Decode packet length & check if data is junk before handling.
		packetLen := protocol.DecodeB64(encodedLen)
		if !validate(s.log, packetLen, reader.Size()) {
			// Reset the reader if we find junk data in the connections buffer.
			reader.Reset(s.connection)
			continue
		}

		// Build a packet object from the remaining bytes.
		packetBz := make([]byte, packetLen)

		if _, err = reader.Read(packetBz); err != nil {
			s.log.Error(
				"Error reading packet data from session",
				"session_address", address,
				"error", err,
			)

			s.DisconnectSession()
			return
		}

		packet, err := packetFromBytes(s.incomingRegistry, packetBz)
		if err != nil {
			s.log.Error(
				"Error composing incoming packet from bytes",
				"session_address", address,
				"packet_header", packet.Header,
				"header_id", packet.HeaderId,
				"payload", packet.Payload.String(),
				"error", err,
			)

			s.DisconnectSession()
			return
		}

		go s.handlePacket(packet, queues)
	}
}

// packetFromBytes attempts to build a packets.IncomingPacket from a slice of bytes.
func packetFromBytes(registry *incoming.IncomingRegistry, packetBytes []byte) (protocol.IncomingPacket, error) {
	payload := bytes.NewBuffer(packetBytes)
	rawHeader := make([]byte, 2)

	for i := 0; i < 2; i++ {
		b, err := payload.ReadByte()
		if err != nil {
			return protocol.IncomingPacket{}, err
		}
		rawHeader[i] = b
	}

	headerID := protocol.DecodeB64(rawHeader)

	info, exists := registry.GetPacketHeader(headerID)
	if !exists {
		// TODO: maybe we should error out here but for now we log the packet info so we can continue
		// to receive data from the client for reversing incoming packet structures.

		return protocol.NewIncoming("", rawHeader, headerID, payload), nil
		//return protocol.NewIncoming("", rawHeader, headerID, payload), fmt.Errorf("packet with header %s not found in IncomingRegistry", string(rawHeader))
	}

	return protocol.NewIncoming(info.Name, rawHeader, headerID, payload), nil
}

// readPacketLength attempts to read the 3 byte Base64 encoded packet length from the buffered reader.
func readPacketLength(reader *bufio.Reader) ([]byte, error) {
	encodedLen := make([]byte, 3)

	for i := 0; i < 3; i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		encodedLen[i] = b
	}

	return encodedLen, nil
}

// validate checks if an incoming packet is valid.
// Its argument is a 3 byte Base64 encoded length.
func validate(log *slog.Logger, packetLen, bytesRead int) bool {
	switch {
	case packetLen == 0:
		log.Debug("Junk packet received")

		return false
	case bytesRead < packetLen:
		log.Debug(
			"Packet length mismatch",
			"expected_length", packetLen,
			"got_length", bytesRead,
		)

		return false
	default:
		return true
	}
}

func (s *Session) handlePacket(packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	handler, exists := s.incomingRegistry.GetCommand(
		protocol.PacketHeader{
			Name:     packet.Name,
			HeaderID: packet.HeaderId,
		},
	)
	if !exists || handler == nil {
		s.log.Debug(
			"Unknown incoming packet",
			"packet_name", packet.Name,
			"session_address", s.address,
			"packet_header", packet.Header,
			"header_id", packet.HeaderId,
			"payload", packet.Payload.String(),
		)

		return
	}

	s.log.Debug(
		"Incoming packet",
		"packet_name", packet.Name,
		"session_address", s.address,
		"packet_header", packet.Header,
		"header_id", packet.HeaderId,
		"payload", packet.Payload.String(),
	)

	handler(s, packet, queues)
}

func (s *Session) ID() string {
	return s.id
}

// Send finalizes an outgoing packet with 0x01 and then attempts to write the packet to a Session's buffer
// before flushing the buffer.
func (s *Session) Send(packet protocol.OutgoingPacket) {
	packet.Finish()

	_, err := s.buff.Write(packet.Payload.Bytes())
	if err != nil {
		s.log.Error(
			"Error writing to session buffer",
			"packet_name", packet.Name,
			"session_address", s.address,
			"packet_header", packet.Header,
			"header_id", packet.HeaderId,
			"payload", packet.Payload.String(),
			"error", err,
		)

		s.DisconnectSession()
		return
	}

	err = s.buff.Flush()
	if err != nil {
		s.log.Error(
			"Error sending packet to session",
			"packet_name", packet.Name,
			"session_address", s.address,
			"packet_header", packet.Header,
			"header_id", packet.HeaderId,
			"payload", packet.Payload.String(),
			"error", err,
		)

		s.DisconnectSession()
		return
	}

	s.log.Debug(
		"Outgoing packet",
		"packet_name", packet.Name,
		"session_address", s.address,
		"packet_header", packet.Header,
		"header_id", packet.HeaderId,
		"payload", packet.Payload.String(),
	)
}

// Queue finalizes an outgoing packet with 0x01 and then attempts to write the packet to a Session's buffer.
func (s *Session) Queue(packet protocol.OutgoingPacket) {
	packet.Finish()

	_, err := s.buff.Write(packet.Payload.Bytes())
	if err != nil {
		s.log.Error(
			"Error writing packet to session buffer",
			"packet_name", packet.Name,
			"session_address", s.address,
			"packet_header", packet.Header,
			"header_id", packet.HeaderId,
			"payload", packet.Payload.String(),
			"error", err,
		)

		s.DisconnectSession()
		return
	}
}

// Flush attempts to flush a Session's buffer.
func (s *Session) Flush() {
	err := s.buff.Flush()
	if err != nil {
		s.log.Error(
			"Error flushing the sessions buffer",
			"session_address", s.address,
			"error", err,
		)

		s.DisconnectSession()
		return
	}
}

// Close closes the underlying network connection to the server and gracefully handles any cleanup.
// This function should only be called from the game server so it can handle the cache of active sessions.
func (s *Session) Close() {
	defer s.connection.Close()

	s.log.Debug("Closing session", "address", s.address)

	// TODO: implement logic to gracefully cleanup a Session.
	// need to prune game state related to the Session from memory and possibly write to disk.
}

// DisconnectSession sends the necessary session information that the server needs to properly handle
// closing an active session over the channel that the server listens to.
func (s *Session) DisconnectSession() {
	s.disconnectedChan <- DisconnectedSession{Address: s.address, SessionID: s.id}
}
