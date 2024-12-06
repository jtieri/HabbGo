package protocol

import (
	"bytes"
)

// OutgoingPacket represents a server->client packet.
type OutgoingPacket struct {
	Name     string
	Header   string
	HeaderId int
	Payload  *bytes.Buffer
}

// NewOutgoing returns a pointer to a newly allocated OutgoingPacket struct.
// The two byte Base64 encoded header is written to the packets buffer on creation for quick composition of packets.
func NewOutgoing(header PacketHeader) OutgoingPacket {
	id := EncodeB64(header.HeaderID, 2)
	packet := OutgoingPacket{
		Name:     header.Name,
		Header:   string(id),
		HeaderId: header.HeaderID,
		Payload:  bytes.NewBuffer(id),
	}
	return packet
}

// Write will write the passed in object as a string
func (packet *OutgoingPacket) Write(s string) {
	packet.Payload.Write([]byte(s))
}

// WriteString writes a string to the packets buffer.
func (packet *OutgoingPacket) WriteString(s string) {
	packet.Payload.Write([]byte(s))
	packet.Payload.WriteByte(2) // FUSEv0.2.0 string parameter ending marker
}

// WriteInt writes a Vl64 encoded int to the packets buffer.
func (packet *OutgoingPacket) WriteInt(i int) {
	packet.Payload.Write(EncodeVl64(i))
}

// WriteBool writes a Vl64 encoded int representing true or false to the packets buffer.
func (packet *OutgoingPacket) WriteBool(b bool) {
	if b {
		packet.WriteInt(1) // H
	} else {
		packet.WriteInt(0) // I
	}
}

// WriteValue writes a key-value entry, separated by '=',to the packets buffer.
func (packet *OutgoingPacket) WriteValue(key []byte, value []byte) {
	packet.Payload.Write(key)
	packet.Payload.Write([]byte("="))
	packet.Payload.Write(value)
	packet.Payload.WriteByte(13) // FUSEv0.2.0 key-value parameter ending marker
}

// WriteKeyValue writes a key-value pair, separated by ':', to the packets buffer.
func (packet *OutgoingPacket) WriteKeyValue(key []byte, value []byte) {
	packet.Payload.Write(key)
	packet.Payload.Write([]byte(":"))
	packet.Payload.Write(value)
	packet.Payload.WriteByte(13) // FUSEv0.2.0 key-value parameter ending marker
}

// WriteDelim writes a custom key-delimeter value to the packets buffer.
func (packet *OutgoingPacket) WriteDelim(key []byte, delim []byte) {
	packet.Payload.Write(key)
	packet.Payload.Write(delim)
}

// String returns the remaining bytes in the packets buffer as a string.
func (packet *OutgoingPacket) String() string {
	return packet.Payload.String()
}

// Finish writes byte 0x01 to the packets buffer to signal the ending of a packet.
func (packet *OutgoingPacket) Finish() {
	packet.Payload.WriteByte(1) // FUSEv0.2.0 server->client packet ending marker
}
