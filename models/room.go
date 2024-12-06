package models

import (
	"time"
)

const (
	WallpaperProperty = "wallpaper"
	FloorProperty     = "floor"

	// heightmapDelimiter is used to mark the end of a line in the text based representation of a Model's Heightmap.
	heightmapDelimiter = "|"

	// PublicRoomOwnerID is used as the owner player ID for rooms in the database.
	PublicRoomOwnerID = 0

	// PublicRoomOffset is used as an offset in the Navigator to easily dichotomize public and private rooms.
	PublicRoomOffset = 1000
)

type Room struct {
	ID              int
	CategoryID      int
	Name            string
	Description     string
	CCTs            string
	Wallpaper       int
	Floor           int
	Landscape       float32
	OwnerID         int
	OwnerName       string
	ShowOwner       bool
	SudoUsers       bool
	Hidden          bool
	AccessType      Access
	Password        string
	CurrentVisitors int
	MaxVisitors     int
	Rating          int
	ChildRooms      []*Room
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Model Model
}

// Model represents a Room's model data.
type Model struct {
	ID        int
	Name      string
	Door      Door
	Heightmap string
}

// Door represents the entrypoint into a Room.
type Door struct {
	X         int
	Y         int
	Z         float64
	Direction int
}
