package commands

import (
	player2 "github.com/jtieri/habbgo/legacy/game/player"
	"github.com/jtieri/habbgo/legacy/protocol/messages"
	"github.com/jtieri/habbgo/legacy/protocol/packets"
)

func INIT_CRYPTO(player player2.Player, packet packets.IncomingPacket) {
	player.Session.Send(messages.CRYPTOPARAMETERS, messages.CRYPTOPARAMETERS())
}

func GENERATEKEY(player player2.Player, packet packets.IncomingPacket) {
	player.Session.Send(messages.AVAILABLESETS, messages.AVAILABLESETS())
	player.Session.Send(messages.ENDCRYPTO, messages.ENDCRYPTO())
	//player.Session.Send(composers.SECRETKEY())
}

func GET_SESSION_PARAMETERS(player player2.Player, packet packets.IncomingPacket) {
	player.Session.Send(messages.SESSIONPARAMETERS, messages.SESSIONPARAMETERS())
}

func VERSIONCHECK(player player2.Player, packet packets.IncomingPacket) {

}

func UNIQUEID(player player2.Player, packet packets.IncomingPacket) {

}

func SECRETKEY(player player2.Player, packets packets.IncomingPacket) {
	player.Session.Send(messages.ENDCRYPTO, messages.ENDCRYPTO())
}

// TODO finish SSO login
func SSO(p player2.Player, packet packets.IncomingPacket) {
	token := packet.ReadString()

	// TODO if p login with token is success login, otherwise send LOCALISED ERROR & disconnect from server
	if token == "" {
		p.Login()
	}
}

func TRY_LOGIN(p player2.Player, packet packets.IncomingPacket) {
	username := packet.ReadString()
	password := packet.ReadString()

	if p.Repo.LoginDB(p, username, password) {
		p.Login()

		// TODO get rid of indirection once NewPlayer returns non-pointer type
		ps := p.Services.PlayerService().(*player2.PlayerServiceProxy)
		ps.AddPlayer(p)

		p.Session.Send(messages.LOGINOK, messages.LOGINOK())
	} else {
		p.Session.Send(messages.LOCALISED_ERROR, messages.LOCALISED_ERROR("Invalid Login Credentials."))
	}
}
