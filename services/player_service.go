package services

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/jtieri/habbgo/models"
	"github.com/jtieri/habbgo/repo/database"
)

type PlayerService struct {
	log *slog.Logger

	queues ServiceQueue
	repo   PlayerRepository

	errChan      chan error
	shutdownChan chan bool

	players map[string]models.Player // Map Session ID to Player
	running bool
}

// NewPlayerService instantiates a new PlayerService and starts it in its own goroutine.
func NewPlayerService(
	ctx context.Context,
	log *slog.Logger,
	errChan chan error,
	shutdownChan chan bool,
	db *sql.DB,
) ServiceQueue {
	ps := PlayerService{
		log:          log,
		queues:       NewServiceQueue(),
		repo:         database.NewPlayerRepo(db),
		errChan:      errChan,
		shutdownChan: shutdownChan,
		players:      make(map[string]models.Player),
		running:      false,
	}

	go ps.start(ctx)

	return ps.queues
}

// start attempts to listen for and handle incoming Message and Request.
// start should only be called once in its own goroutine.
func (ps *PlayerService) start(ctx context.Context) {
	if ps.running {
		ps.errChan <- errors.New("the PlayerService is already running")
		return
	}

	ps.running = true

	defer func() {
		ps.running = false
	}()

	ps.log.Debug("PlayerService is running")

	for {
		select {
		case msg, open := <-ps.queues.msgQueue:
			if !open {
				ps.log.Debug("PlayerService msg queue is closed")
				return
			}

			ps.log.Debug("PlayerService received a message from the queue")

			msg.Handle(ps)
		case req, open := <-ps.queues.reqQueue:
			if !open {
				ps.log.Debug("PlayerService req queue is closed")
				return
			}

			ps.log.Debug("PlayerService received a request from the queue")

			req.Respond(ps)
		case <-ctx.Done():
			ps.log.Debug("Context cancelled")
			ps.log.Debug("Shutting down PlayerService...")

			ps.shutdownChan <- true

			return
		}
	}
}

func (ps *PlayerService) LoginPlayer(sessionID, username, password string) bool {
	// TODO: make a database call to retrieve player for specified username and password
	/*
		player, ok := ps.repo.Player(sessionID, username, password)
		if !ok {
			return false
		}

		ps.players[sessionID] = p

		return true
	*/

	p := models.Player{
		ID:           1234,
		Username:     username,
		Figure:       "",
		Sex:          "M",
		Motto:        "habbgo rocks!",
		ConsoleMotto: "",
		Tickets:      10,
		PoolFigure:   "",
		Film:         20,
		Credits:      30,
		LastOnline:   time.Now(),
		PlayerRank:   0,
		Badges:       []string{},
		CurrentBadge: "",
		DisplayBadge: false,
		SoundEnabled: false,
		LoggedIn:     true,
	}

	ps.players[sessionID] = p

	return true
}

func (ps *PlayerService) CreditBalance(sessionID string) int {
	player, ok := ps.players[sessionID]
	if !ok {
		// TODO: sign of a critical issue, handle gracefully
		ps.logPlayerNotFound(sessionID)
	}

	return player.Credits
}

func (ps *PlayerService) SoundSetting(sessionID string) bool {
	player, ok := ps.players[sessionID]
	if !ok {
		// TODO: sign of a critical issue, handle gracefully
		ps.logPlayerNotFound(sessionID)
	}

	return player.SoundEnabled
}

type BadgeDetails struct {
	Badges       []string
	CurrentBadge string
	DisplayBadge bool
}

func (ps *PlayerService) BadgeInfo(sessionID string) BadgeDetails {
	player, ok := ps.players[sessionID]
	if !ok {
		// TODO: sign of a critical issue, handle gracefully
		ps.logPlayerNotFound(sessionID)
	}

	return BadgeDetails{
		Badges:       player.Badges,
		CurrentBadge: player.CurrentBadge,
		DisplayBadge: player.DisplayBadge,
	}
}

type UserObjectInfo struct {
	ID         int
	Username   string
	Figure     string
	PoolFigure string
	Sex        string
	Motto      string
	Tickets    int
	Film       int
}

func (ps *PlayerService) UserObjInfo(sessionID string) UserObjectInfo {
	player, ok := ps.players[sessionID]
	if !ok {
		// TODO: sign of a critical issue, handle gracefully
		ps.logPlayerNotFound(sessionID)
	}

	return UserObjectInfo{
		ID:         player.ID,
		Username:   player.Username,
		Figure:     player.Figure,
		PoolFigure: player.PoolFigure,
		Sex:        player.Sex,
		Motto:      player.Motto,
		Tickets:    player.Tickets,
		Film:       player.Film,
	}
}

func (ps *PlayerService) PlayerRank(sessionID string) models.Rank {
	player, ok := ps.players[sessionID]
	if !ok {
		// TODO: sign of a critical issue, handle gracefully
		ps.logPlayerNotFound(sessionID)
	}

	return player.PlayerRank
}

func (ps *PlayerService) logPlayerNotFound(sessionID string) {
	ps.log.Error("Player not found in PlayerService cache", "session_id", sessionID)
}
