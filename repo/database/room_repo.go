package database

import (
	"database/sql"
	"strings"

	"github.com/jtieri/habbgo/models"
)

type RoomRepo struct {
	database *sql.DB
}

func NewRoomRepo(db *sql.DB) *RoomRepo {
	return &RoomRepo{database: db}
}

func (rr *RoomRepo) PublicRooms() ([]models.Room, error) {
	stmt, err := rr.database.Prepare(
		"SELECT r.*, rm.* FROM rooms r LEFT JOIN room_categories rc ON r.category_id = rc.id LEFT JOIN room_models rm on r.model_id = rm.id WHERE rc.is_public=true",
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		room, err := fillRoomData(rows)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

// TODO: implement me
func (rr *RoomRepo) RoomsByPlayerID(playerID int) ([]models.Room, error) {
	return nil, nil
}

// TODO: implement me
func (rr *RoomRepo) RoomByID(playerID int) (models.Room, error) {
	return models.Room{}, nil
}

func fillRoomData(rows *sql.Rows) (models.Room, error) {
	var tmpAccessType string

	room := models.Room{}

	err := rows.Scan(
		&room.ID,
		&room.CategoryID,
		&room.Name,
		&room.Description,
		&room.OwnerID,
		&room.Model.ID,
		&room.CCTs,
		&room.Wallpaper,
		&room.Floor,
		&room.ShowOwner,
		&room.Password,
		&tmpAccessType,
		&room.SudoUsers,
		&room.CurrentVisitors,
		&room.MaxVisitors,
		&room.Rating,
		&room.Hidden,
		&room.CreatedAt,
		&room.UpdatedAt,
		&room.Model.ID,
		&room.Model.Name,
		&room.Model.Door.X,
		&room.Model.Door.Y,
		&room.Model.Door.Z,
		&room.Model.Door.Direction,
		&room.Model.Heightmap,
	)
	if err != nil {
		return models.Room{}, err
	}

	// find the right Access for the string representation of an Access
	// that we loaded from the database.
	for _, access := range models.AccessTypes() {
		if strings.ToLower(access.String()) == tmpAccessType {
			room.AccessType = access
		}
	}

	// TODO: build Room height map from database
	// When rooms are loaded from the database we want to be sure that we are building their map
	// from the room model's heightmap.
	//room, err = parseHeightMap(room)
	//if err != nil {
	//	return models.Room{}, err
	//}

	return room, nil
}
