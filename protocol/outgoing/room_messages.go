package outgoing

import "github.com/jtieri/habbgo/protocol"

// InterstitialData is sent from the server as a response to commands.GETINTERST.
// It's payload can contain two strings where one is a URL to an image to be used
// during the loading room UI and the second one is a URL to a destination users will go if the image is clicked.
func InterstitialData() protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.InterstitialData)

	/*
		This is the client lingo code for handling this message.
		We should be checking if there is an interstitial set for loading and if so send
		the URL to the image and the URL to the destination site.

		if tMsg.content.length > 1 then
		    tDelim = the itemDelimiter
		    the itemDelimiter = "\t"
		    tSourceURL = tMsg.content.getProp(#item, 1)
		    tTargetURL = tMsg.content.getProp(#item, 2)
		    the itemDelimiter = tDelim
		    me.getComponent().getInterstitial().Init(tSourceURL, tTargetURL)
	*/

	return packet
}

// OpenConnectionOk is sent from the server as a response to commands.ROOM_DIRECTORY.
// It will confirm to the client that the room the player is currently trying to enter is
// actually a private room.
func OpenConnectionOk() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.OpcOk)
}

// RoomReady is sent from the server as one of many packets in response to commands.ROOM_DIRECTORY.
// It sends the appropriate room ID and room model name to the client for initializing the room.
func RoomReady() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.RoomReady)
}

// FlatProperty is sent from the server as one of the many packets in response to commands.ROOM_DIRECTORY.
// It contains a room property and value which is used in updating the wallpaper and floor elements
// in the room.
func FlatProperty() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.FlatProperty)
}

// RoomRating is sent from the server to update a rooms ratings in the client.
func RoomRating() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.RoomRating)
}

// RoomAd is sent from the server as a response to commands.GETROOMAD.
// It will send the rooms ad image URL and the target URL for where users end up if they click the ad.
func RoomAd() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.RoomAd)
}

// HeightMap is sent from the server as a response to commands.G_HMAP.
// It sends the room's heightmap as a string to the client to for rendering.
func HeightMap() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.Heightmap)
}

// UserObjects is sent from the server as a response to commands.G_USRS.
// It serializes the state for each entity in the current room and sends the data to the client for rendering.
func UserObjects() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.Users)
}

// Objects is sent from the server as a response to commands.G_OBJS.
// It serializes the public room items in the current room and sends them to the client for rendering.
func Objects() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.Objects)
}

// ActiveObjects is sent from the server as a response to commands.G_OBJS.
// It serializes the floor items in the current room and sends them to the client for rendering.
func ActiveObjects() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.ActiveObjects)
}

func Items() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.Items45)
}

func Status() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.Status)
}

func HotelView() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.Clc)
}

func serializeItems() {

}
