package outgoing

import "github.com/jtieri/habbgo/protocol"

func Date(date string) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.Date)
	packet.Write(date)
	return packet
}

func ApproveNameReply(code int) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.ApproveNameReply)
	packet.WriteInt(code)
	return packet
}

func NameUnacceptable(code int) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.NameUnacceptable)
	packet.WriteInt(0)
	return packet
}

func PasswordApproved(code int) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.PasswordApproved)
	packet.WriteInt(code)
	return packet
}

func EmailApproved() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.EmailApproved)
}

func EmailRejected() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.EmailRejected)
}
