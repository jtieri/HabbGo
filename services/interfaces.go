package services

import (
	"context"

	"github.com/jtieri/habbgo/models"
)

// Messager should be implemented by messages in other packages that need to execute some logic against
// cached state managed by a service.
type Messager interface {
	Handle(s Servicer)
}

// Requester should be implemented by requests in other packages that need to query some cached state managed by a service.
type Requester interface {
	Respond(s Servicer)
}

// Servicer should be implemented by the various services so that they can be used in Messager & Requester implementations.
type Servicer interface {
	start(ctx context.Context)
}

// PlayerRepository should be implemented by data access types to retrieve Player information from
// some underlying data store.
type PlayerRepository interface {
	Register(username, figure, gender, email, birthday, createdAt, password string, salt []byte) error
	Login(username, password string) (models.Player, error)
	LoadBadges(playerID int) []string
	PlayerExists(username string) bool
}

// RoomRepository should be implemented by data access types to retrieve Room information from
// some underlying data store.
type RoomRepository interface {
	PublicRooms() ([]models.Room, error)
	RoomsByPlayerID(playerID int) ([]models.Room, error)
	RoomByID(roomID int) (models.Room, error)
}

// NavigatorRepository should be implemented by data access types to retrieve Category information from
// some underlying data store.
type NavigatorRepository interface {
	Categories() ([]models.Category, error)
}
