package outgoing

import (
	"strconv"
	"strings"

	"github.com/jtieri/habbgo/models"
	"github.com/jtieri/habbgo/protocol"
)

func NavNodeInfo(
	parentCat models.Category,
	subcategories []SubCategoryInfo,
	rooms []models.Room,
	hideFullRooms bool,
	currentVisitors int,
	maxVisitors int,
	username string,
	playerRank models.Rank,
) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.NavNodeInfo)

	packet.WriteBool(hideFullRooms) // hideCategory
	packet.WriteInt(parentCat.ID)

	if parentCat.IsPublic {
		packet.WriteInt(0)
	} else {
		packet.WriteInt(2)
	}

	packet.WriteString(parentCat.Name)
	packet.WriteInt(currentVisitors)
	packet.WriteInt(maxVisitors)
	packet.WriteInt(parentCat.ParentID)

	if !parentCat.IsPublic {
		packet.WriteInt(len(rooms))
	}

	for _, room := range rooms {
		if room.OwnerID == models.PublicRoomOwnerID {
			desc := room.Description

			var door int
			if strings.Contains(desc, "/") {
				data := strings.Split(desc, "/")
				desc = data[0]
				door, _ = strconv.Atoi(data[1])
			}

			packet.WriteInt(room.ID + models.PublicRoomOffset) // writeInt roomId
			packet.WriteInt(1)                                 // writeInt 1
			packet.WriteString(room.Name)                      // writeString roomName
			packet.WriteInt(room.CurrentVisitors)              // writeInt currentVisitors
			packet.WriteInt(room.MaxVisitors)                  // writeInt maxVisitors
			packet.WriteInt(room.CategoryID)                   // writeInt catId
			packet.WriteString(desc)                           // writeString roomDesc
			packet.WriteInt(room.ID)                           // writeInt roomId
			packet.WriteInt(door)                              // writeInt door
			packet.WriteString(room.CCTs)                      // writeString roomCCTs
			packet.WriteInt(0)                                 // writeInt 0
			packet.WriteInt(1)                                 // writeInt 1
		} else {
			packet.WriteInt(room.ID)
			packet.WriteString(room.Name)

			// TODO check that player is owner of r, that r is showing owner name, or that player has right SEE_ALL_ROOMOWNERS
			if username == room.OwnerName || room.ShowOwner {
				packet.WriteString(room.OwnerName)
			} else {
				packet.WriteString("-")
			}

			packet.WriteString(room.AccessType.String())
			packet.WriteInt(room.CurrentVisitors)
			packet.WriteInt(room.MaxVisitors)
			packet.WriteString(room.Description)
		}
	}

	// iterate over sub-categories
	for _, subcat := range subcategories {
		if subcat.MinRankAccess > playerRank {
			continue
		}

		packet.WriteInt(subcat.ID)
		packet.WriteInt(0)
		packet.WriteString(subcat.Name)
		packet.WriteInt(subcat.CurrentVisitors) // writeInt currentVisitors
		packet.WriteInt(subcat.MaxVisitors)     // writeInt maxVisitors
		packet.WriteInt(parentCat.ID)
	}

	return packet
}

type SubCategoryInfo struct {
	ID              int
	Name            string
	CurrentVisitors int
	MaxVisitors     int
	MinRankAccess   models.Rank
}

// UserFlatCats is sent from the server in response to GETUSERFLATCATS.
// It contains the Navigator category information for private rooms that should
// be visible for the specified player.
func UserFlatCats(categories []models.Category) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.UserFlatCats)

	packet.WriteInt(len(categories))

	for _, cat := range categories {
		packet.WriteInt(cat.ID)
		packet.WriteString(cat.Name)
	}

	return packet
}
