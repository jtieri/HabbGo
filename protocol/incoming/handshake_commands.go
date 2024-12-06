package incoming

import (
	"fmt"

	"github.com/jtieri/habbgo/protocol"
	"github.com/jtieri/habbgo/protocol/outgoing"
	"github.com/jtieri/habbgo/services"
	"github.com/jtieri/habbgo/services/requests"
)

type Session interface {
	ID() string
	Send(packet protocol.OutgoingPacket)
	Queue(packet protocol.OutgoingPacket)
	Flush()
}

func InitCrypto(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	session.Send(outgoing.CryptoParameters())
}

func GenerateKey(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	session.Send(outgoing.AvailableSets())
	session.Send(outgoing.EndCrypto())
}

func GetSessionParameters(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	session.Send(outgoing.SessionParameters())
}

// VersionCheck handles the VERSIONCHECK packet from the client.
// It's payload contains:
// - Base64 encoded header @E
// - VL64 encoded client version
// - Client URL, when using a projector this value ends up being the sw variable.
// - URL to the external_variables.txt file, when using a projector this value may reference a file in the filesystem.
// Ex: @EYdA@A2@Vexternal_variables.txt
// Where @E is the header, YdA is the client version, @A is the client URL string length, 2 is the client URL,
// @V is the external_variables.txt URL string length, external_variables.txt is the external variables URL.
func VersionCheck(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	clientVersion := packet.ReadInt()
	clientURL := packet.ReadString()
	extVarsURL := packet.ReadString()

	fmt.Printf("Client version: %d \n", clientVersion)
	fmt.Printf("Client URL: %s \n", clientURL)
	fmt.Printf("External vars URL: %s \n", extVarsURL)
}

// UniqueID handles the UNIQUEID packet from the client.
// It's payload contains:
// - Base64 encoded header @F
// - 2 byte Base64 encoded length of the payload string
// - The machine ID of the connected machine, retrieved via getMachineID in the client code.
// Ex: @F@S2754237384254102022
// Where @F is the header, @S is the string length, 2754237384254102022 is the string.
func UniqueID(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	machineID := packet.ReadString()
	_ = machineID

	/*
		on getMachineID
		  return(getSpecialServices().getMachineID())
		end
	*/

	// TODO: persist machine IDs to the database to use for managing bans at the machine level.
}

func SecretKey(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	session.Send(outgoing.EndCrypto())
}

func SSO(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	token := packet.ReadString()

	// TODO: if player with specified token login is successful login, otherwise send localised error & disconnect
	if token == "" {

	}
}

func TryLogin(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	username := packet.ReadString()
	password := packet.ReadString()

	req := &requests.LoginPlayerReq{
		SessionID: session.ID(),
		Username:  username,
		Password:  password,
		Response:  make(chan bool),
	}

	queues.PlayerQueues.QueueRequest(req)

	loginOk := <-req.Response

	if loginOk {
		session.Send(outgoing.LoginOK())
	} else {
		session.Send(outgoing.LocalisedError("Wrong username or password - please try again!"))
	}
}
