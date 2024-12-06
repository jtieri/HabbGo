package commands

import (
	"github.com/jtieri/habbgo/legacy/game/player"
	messages2 "github.com/jtieri/habbgo/legacy/protocol/messages"
	"github.com/jtieri/habbgo/legacy/protocol/packets"
)

func GET_INFO(player player.Player, packet packets.IncomingPacket) {
	player.Session.Send(messages2.USEROBJ, messages2.USEROBJ(player))
}

func GET_CREDITS(player player.Player, packet packets.IncomingPacket) {
	player.Session.Send(messages2.CREDITBALANCE, messages2.CREDITBALANCE(player.Details.Credits))
}

func GETAVAILABLEBADGES(player player.Player, packet packets.IncomingPacket) {
	player.Session.Send(messages2.AVAILABLESETS, messages2.AVAILABLEBADGES(player))
}

func GET_SOUND_SETTING(player player.Player, packet packets.IncomingPacket) {
	player.Session.Send(messages2.SOUNDSETTING, messages2.SOUNDSETTING(player.Details.SoundEnabled))
}

func TestLatency(player player.Player, packet packets.IncomingPacket) {
	l := packet.ReadInt()
	player.Session.Send(messages2.Latency, messages2.Latency(l))
}
