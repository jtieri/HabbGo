package incoming

import (
	"fmt"

	"github.com/jtieri/habbgo/protocol"
	"github.com/jtieri/habbgo/protocol/outgoing"
	"github.com/jtieri/habbgo/services"
)

// GetInterst is sent from the client when attempting to enter a room from the Navigator.
// It attempts to get interstitial data used when displaying the room loading UI.
func GetInterst(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	session.Send(outgoing.InterstitialData())
}

// RoomDirectory is sent from the client when trying to enter a room.
// It will determine if the room is public or private and perform the necessary
// checks before initiating the server side logic for entering a room and,
// sending the response packets.
func RoomDirectory(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	isPublicRoom := packet.ReadBytes(1)
	roomID := packet.ReadInt()
	doorID := packet.ReadInt()

	_ = roomID
	_ = doorID

	fmt.Printf("IsPublicRoom: %s \n", string(isPublicRoom))
	fmt.Printf("Room ID: %d \n", roomID)
	fmt.Printf("Door ID: %d \n", doorID)

	// Send open connection ok for private rooms
	// A is 1 Base64 encoded
	if string(isPublicRoom) != "A" {
		session.Send(outgoing.OpenConnectionOk())
		return
	}

}

func GetRoomAd(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	session.Send(outgoing.RoomAd())
}

func GHmap(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	session.Send(outgoing.HeightMap())
}

func GUsrs(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {

}

func GObjs(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {

}

func GStat(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {

}

func GoAway(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {

}

func Move(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {

}

func Stop(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {

}

func Quit(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	// TODO: update player state to leave room

	// TODO: if user is in room, send packet LOGOUT to every user in the associated room
	// to ensure state is updated to remove the player from the room

	// TODO: this is probably wrong, using this for now to go to hotel view on cancelling room entry in UI
	// before a room is properly initialized

	session.Send(outgoing.HotelView())
}
