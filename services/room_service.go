package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jtieri/habbgo/models"
	"github.com/jtieri/habbgo/repo/database"
)

type RoomService struct {
	log *slog.Logger

	queues ServiceQueue
	repo   RoomRepository

	errChan      chan error
	shutdownChan chan bool

	rooms   map[int]models.Room
	running bool
}

// NewRoomService instantiates a new RoomService and starts it in its own goroutine.
func NewRoomService(
	ctx context.Context,
	log *slog.Logger,
	errChan chan error,
	shutdownChan chan bool,
	db *sql.DB,
) ServiceQueue {
	rs := RoomService{
		log:          log,
		queues:       NewServiceQueue(),
		repo:         database.NewRoomRepo(db),
		errChan:      errChan,
		shutdownChan: shutdownChan,
		rooms:        make(map[int]models.Room),
		running:      false,
	}

	go rs.start(ctx)

	return rs.queues
}

// start attempts to listen for and handle incoming Message and Request.
// start should only be called once in its own goroutine.
func (rs *RoomService) start(ctx context.Context) {
	if rs.running {
		rs.errChan <- errors.New("the RoomService is already running")
		return
	}

	//publicRooms, err := rs.repo.LoadPublicRooms()
	//if err != nil {
	//	// TODO: handle database related errors. This probably should result in a retry or the root ctx being cancelled
	//	// by the game server.
	//
	//	rs.log.Error("Failed to get public rooms from database", "error", err)
	//}
	//
	//for _, room := range publicRooms {
	//	rs.rooms[room.ID] = room
	//}

	rs.running = true

	defer func() {
		rs.running = false
	}()

	rs.log.Debug("RoomService is running")

	for {
		select {
		case msg, open := <-rs.queues.msgQueue:
			if !open {
				rs.log.Debug("RoomService msg queue is closed")
				return
			}

			rs.log.Debug("RoomService received a message from the queue")

			msg.Handle(rs)
		case req, open := <-rs.queues.reqQueue:
			if !open {
				rs.log.Debug("RoomService req queue is closed")
				return
			}

			rs.log.Debug("RoomService Received a request from the queue")

			req.Respond(rs)
		case <-ctx.Done():
			rs.log.Debug("Context cancelled")
			rs.log.Debug("Shutting down RoomService...")

			rs.shutdownChan <- true

			return
		}
	}
}

func (rs *RoomService) RoomByID(id int) models.Room {
	room, ok := rs.rooms[id]
	if ok {
		return room
	}

	room, err := rs.repo.RoomByID(id)
	if err != nil {
		// TODO: handle database related errors.
		// This probably should result in a retry or the root ctx being cancelled by the game server.

		rs.log.Error("Failed to get room from database", "room_id", id, "error", err)
		return models.Room{}
	}

	rs.rooms[id] = room
	return room
}

func (rs *RoomService) PublicRooms() []models.Room {
	var rooms []models.Room

	publicRooms, err := rs.repo.PublicRooms()
	if err != nil {
		// TODO: handle database related errors.
		// This probably should result in a retry or the root ctx being cancelled by the game server.

		rs.log.Error("Failed to get public rooms from database", "error", err)
		return nil
	}

	// For each room if the room is already cached use the cached version.
	for _, room := range publicRooms {
		if cachedRoom, ok := rs.rooms[room.ID]; ok {
			rooms = append(rooms, cachedRoom)
		} else {
			rooms = append(rooms, room)
		}
	}

	fmt.Printf("Number of public rooms in rs.PublicRooms: %d\n", len(rooms))

	return rooms
}

func (rs *RoomService) CurrentVisitorsForCategory(categoryID int) int {
	visitors := 0

	for _, room := range rs.rooms {
		if room.CategoryID == categoryID {
			visitors += room.CurrentVisitors
		}
	}

	return visitors
}

func (rs *RoomService) MaxVisitorsForCategory(categoryID int) int {
	visitors := 0

	for _, room := range rs.rooms {
		if room.CategoryID == categoryID {
			visitors += room.MaxVisitors
		}
	}

	return visitors
}

func (rs *RoomService) MaxVisitorsForSubCategory(categoryID, roomID int) int {
	room, ok := rs.rooms[roomID]
	if !ok {
		// TODO: sign of a bigger issue at hand, handle this gracefully
	}

	if room.CategoryID == categoryID {
		return room.MaxVisitors
	}

	return -1
}
