package outgoing

import "github.com/jtieri/habbgo/protocol"

type OutgoingRegistry struct {
	RegisteredPackets map[string]protocol.PacketHeader
}

func NewOutgoingRegistry() *OutgoingRegistry {
	return &OutgoingRegistry{
		RegisteredPackets: make(map[string]protocol.PacketHeader),
	}
}

func (ir *OutgoingRegistry) RegisterOutgoingHeadersv14() {
	// Handshake packet headers
	ir.RegisteredPackets["@@"] = protocol.Hello
	ir.RegisteredPackets["DU"] = protocol.CryptoParameters
	ir.RegisteredPackets["@A"] = protocol.SecretKeyOutgoing
	ir.RegisteredPackets["DV"] = protocol.EndCrypto
	ir.RegisteredPackets["DA"] = protocol.SessionParameters
	ir.RegisteredPackets["@H"] = protocol.AvailableSets
	ir.RegisteredPackets["@C"] = protocol.LoginOk
	ir.RegisteredPackets["@a"] = protocol.LocalisedError

	// Registration packet headers
	ir.RegisteredPackets["Bc"] = protocol.Date
	ir.RegisteredPackets["@d"] = protocol.ApproveNameReply
	ir.RegisteredPackets["@e"] = protocol.NameUnacceptable
	ir.RegisteredPackets["DZ"] = protocol.PasswordApproved
	ir.RegisteredPackets["DO"] = protocol.EmailApproved
	ir.RegisteredPackets["DP"] = protocol.EmailRejected
}
