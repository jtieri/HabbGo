package outgoing

import (
	"strconv"

	"github.com/jtieri/habbgo/protocol"
	"github.com/jtieri/habbgo/services"
)

func UserObj(info services.UserObjectInfo) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.UserObj)

	packet.WriteString(strconv.Itoa(info.ID))
	packet.WriteString(info.Username)
	packet.WriteString(info.Figure)
	packet.WriteString(info.Sex)
	packet.WriteString(info.Motto)
	packet.WriteInt(info.Tickets)
	packet.WriteString(info.PoolFigure)
	packet.WriteInt(info.Film)

	// TODO: persist the receive direct mail configuration to the Player model and use that instead of hard coded value.
	directMail := 1
	packet.WriteInt(directMail)

	return packet
}

func CreditBalance(balance int) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.CreditBalance)
	packet.WriteString(strconv.Itoa(balance) + ".0")
	return packet
}

func AvailableBadges(details services.BadgeDetails) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.AvailableBadges)

	packet.WriteInt(len(details.Badges))

	var bSlot int
	for slot, badge := range details.Badges {
		packet.WriteString(badge)

		if badge == details.CurrentBadge {
			bSlot = slot
		}
	}

	packet.WriteInt(bSlot)
	packet.WriteBool(details.DisplayBadge)

	return packet
}

func SoundSetting(enabled bool) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.SoundSetting)
	packet.WriteBool(enabled)
	return packet
}

func Latency(l int) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.Latency)
	packet.WriteInt(l)
	return packet
}
