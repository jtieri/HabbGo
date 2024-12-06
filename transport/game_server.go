package transport

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jtieri/habbgo/config"
	"github.com/jtieri/habbgo/services"
)

// GameServer represents the primary TCP server that handles incoming user connections.
type GameServer struct {
	log      *slog.Logger
	cfg      *config.Config
	database *sql.DB
	sessions map[string][]*Session
}

// DisconnectedSession represents the information necessary for handling disconnected Sessions.
// Sessions may need to be disconnected due to underlying network errors, the hotel closing, or proper logout.
type DisconnectedSession struct {
	Address   string
	SessionID string
}

func NewGameServer(log *slog.Logger, cfg *config.Config, db *sql.DB) *GameServer {
	return &GameServer{
		log:      log,
		cfg:      cfg,
		database: db,
		sessions: make(map[string][]*Session),
	}
}

func (s *GameServer) Start(ctx context.Context) chan error {
	s.log.Info("Starting game server", "host", s.cfg.Server.Host, "port", s.cfg.Server.Port)
	errChan := make(chan error)
	go s.handleConnections(ctx, errChan)
	return errChan
}

func (s *GameServer) handleConnections(ctx context.Context, errChan chan error) {
	disconnectedSessionChan := make(chan DisconnectedSession)

	serviceErrChan := make(chan error)
	serviceShutdownChan := make(chan bool)

	defer close(serviceShutdownChan)
	defer close(serviceErrChan)
	defer close(disconnectedSessionChan)
	defer close(errChan)

	address := fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port)

	localAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		errChan <- err
		return
	}

	// Use ListenTCP vs. net.Listen so that we can set a deadline on the listener,
	listener, err := net.ListenTCP("tcp", localAddr)
	if err != nil {
		errChan <- err
		return
	}
	defer listener.Close()

	s.log.Info("Successfully started the game server", "address", listener.Addr().String())

	queues := s.startGameServices(ctx, serviceErrChan, serviceShutdownChan)

	for {
		select {
		case <-ctx.Done():
			// Main context was cancelled, call shutdown to close sessions and wait for services to shut down.
			//s.log.Debug("Context was cancelled and acked in GameServer.handleConnections")
			s.shutdown(queues, serviceShutdownChan)

			return
		case info := <-disconnectedSessionChan:
			// A Session encountered an internal error and should be closed and removed from the cached Sessions.
			s.handleDisconnectedSession(info)
		case err := <-serviceErrChan:
			// TODO: properly check if error is recoverable or should result in shut down
			s.log.Error("Server received an error from a service", "error", err)
		default:
			// Set a deadline so that we don't stay blocking forever during listener.Accept(),
			// this allows us to gracefully shutdown if the context is cancelled.
			if err := listener.SetDeadline(time.Now().Add(time.Second)); err != nil {
				continue
			}

			// Block and listen for incoming connections.
			conn, err := listener.Accept()
			if err != nil {
				if os.IsTimeout(err) {
					continue
				}

				s.log.Warn("Error trying to handle incoming connection", "error", err)

				continue
			}

			connectionAddress := getAddress(conn.LocalAddr().String())

			// Check if there are already too many active sessions for this address.
			s.checkForActiveSessions(connectionAddress)

			// Create a new Session object and start listening for incoming data over the wire.
			s.createNewSession(ctx, conn, connectionAddress, disconnectedSessionChan, queues)

			s.log.Debug("Accepted new game session", "address", connectionAddress)
		}
	}
}

// Shutdown closes all active Sessions and gracefully cleans up the server.
func (s *GameServer) shutdown(queues *services.ServiceQueues, serviceShutdownChan chan bool) {
	s.log.Info("Shutting down game server...")

	for _, sessions := range s.sessions {
		for _, session := range sessions {
			session.Close()
		}
	}

	count := 0
	for {
		select {
		case <-serviceShutdownChan:
			count += 1

			if count == queues.NumberOfServices {
				s.log.Info("Game services have all shut down successfully.")
				return
			}
		case <-time.After(time.Minute * 1):
			s.log.Error("Failed to properly shut down game services after one minute")
			return
		}
	}
}

func (s *GameServer) startGameServices(
	ctx context.Context,
	serviceErrChan chan error,
	serviceShutdownChan chan bool,
) *services.ServiceQueues {
	s.log.Info("Starting game services...")

	psQueues := services.NewPlayerService(ctx, s.log, serviceErrChan, serviceShutdownChan, s.database)
	navQueues := services.NewNavigatorService(ctx, s.log, serviceErrChan, serviceShutdownChan, s.database)
	rsQueues := services.NewRoomService(ctx, s.log, serviceErrChan, serviceShutdownChan, s.database)

	queues := services.NewServiceQueues(psQueues, navQueues, rsQueues)

	s.log.Info("Successfully started game services")

	return queues
}

// handleDisconnectedSession will close an active session for some session that has reported an error.
// It will also update the map of active sessions appropriately.
func (s *GameServer) handleDisconnectedSession(info DisconnectedSession) {
	// A Session disconnected due to an underlying error.
	sessions, exist := s.sessions[info.Address]
	if !exist {
		// If we have no cached Sessions for this address it is a sign of a fatal error.
		s.log.Error("Disconnected session does not exist", "address", info.Address)

		// TODO: gracefully handle this issue
		return
	}

	// Close the Session and update the cached Sessions appropriately.
	for i, session := range sessions {
		if session.id != info.SessionID {
			continue
		}

		session.Close()

		sessions = slices.Delete(sessions, i, i+1)

		s.sessions[info.Address] = sessions

		if len(s.sessions) == 0 {
			delete(s.sessions, info.Address)
		}
	}
}

// checkForActiveSessions checks if there are already active Sessions for the remote address.
// If there are already more active connections than the configured max amount,
// the first active Session in the cache will be closed and removed from the cache.
func (s *GameServer) checkForActiveSessions(address string) {
	// Check if there are already too many active sessions for this address.
	if s.sessionsFromSameAddr(address) >= s.cfg.Server.MaxConnsPerPlayer {
		s.log.Debug("Too many sessions already connected for this address", "address", address)

		// Disconnect the already existing session for this address and remove it from the slice
		// of active sessions.
		sessions, _ := s.sessions[address]
		session := sessions[0]

		s.log.Debug("Closing session", "address", address, "session_id", session.id)

		session.Close()
		sessions = slices.Delete(sessions, 0, len(sessions)-1)
		s.sessions[address] = sessions
	}
}

// createNewSession will create a new Session object and start its listener for listening for incoming data over the wire.
func (s *GameServer) createNewSession(
	ctx context.Context,
	conn net.Conn,
	address string,
	disconnectedSessionChan chan DisconnectedSession,
	serviceQueues *services.ServiceQueues,
) {
	// Generate a unique user ID for the Session.
	sessionID := uuid.New()

	session := NewSession(s.log, conn, address, sessionID.String(), disconnectedSessionChan)

	// Properly add the new Session to the cache of active Sessions.
	sessions, exists := s.sessions[address]
	if exists {
		sessions = append(sessions, session)
	} else {
		s.sessions[address] = make([]*Session, 1)
		s.sessions[address][0] = session
	}

	go session.listen(ctx, serviceQueues)
}

// address returns the IP address of a Session's connection by splitting the socket address,
// e.g. given the input 127.0.0.1:1234 the function would return 127.0.0.1.
func getAddress(remoteAddress string) string {
	return strings.Split(remoteAddress, ":")[0]
}

// sessionsFromSameAddr returns the number of active Sessions connected to the server for a specified IP address.
func (s *GameServer) sessionsFromSameAddr(address string) int {
	count := 0

	sessions, exists := s.sessions[address]
	if !exists {
		return count
	}

	for _, session := range sessions {
		if address == session.address {
			count++
		}
	}

	return count
}
