package incoming

import (
	"net/mail"
	"strings"
	"time"

	"github.com/jtieri/habbgo/legacy/crypto"
	"github.com/jtieri/habbgo/legacy/date"
	"github.com/jtieri/habbgo/legacy/text"
	"github.com/jtieri/habbgo/protocol"
	"github.com/jtieri/habbgo/protocol/outgoing"
	"github.com/jtieri/habbgo/services"
)

const (
	NameOk              = 0
	NameTooLong         = 1
	NameTooShort        = 2
	NameUnacceptable    = 3
	NameAlreadyReserved = 4

	PasswordOk            = 0
	PasswordTooShort      = 1
	PasswordTooLong       = 2
	PasswordUnacceptable  = 3
	PasswordHasNoNum      = 4
	PasswordSimilarToName = 5
)

func GetAvailableSets(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	session.Send(outgoing.AvailableSets())
}

func GDate(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	session.Send(outgoing.Date(date.CurrentDate()))
}

func ApproveName(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	name := text.Filter(packet.ReadString())
	session.Send(outgoing.ApproveNameReply(checkName(name)))
}

func ApprovePassword(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	username := packet.ReadString()
	password := packet.ReadString()
	session.Send(outgoing.PasswordApproved(checkPassword(username, password)))
}

func ApproveEmail(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	email := packet.ReadString()

	if _, err := mail.ParseAddress(email); err != nil {
		session.Send(outgoing.EmailRejected())
	} else {
		session.Send(outgoing.EmailApproved())
	}
}

func Register(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	packet.ReadB64()
	username := packet.ReadString()

	packet.ReadB64()
	figure := packet.ReadString()

	packet.ReadB64()
	gender := packet.ReadString()

	packet.ReadB64()
	packet.ReadB64()

	packet.ReadB64()
	email := packet.ReadString()

	packet.ReadB64()
	birthday := packet.ReadString()

	packet.ReadBytes(11)
	password := packet.ReadString()

	// hash password before storing in db
	randSalt := crypto.GenerateRandomSalt(crypto.SaltSize)
	hashedPassword := crypto.HashPassword(password, randSalt)

	// generate date time stamp in UTC with format YYYY-MM-DD T HH-MM-SS
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	createdAt := now.Format("2006-01-02 15:04:05")

	_ = createdAt      // TODO: delete
	_ = hashedPassword // TODO: delete
	_ = username       // TODO: delete
	_ = figure         // TODO: delete
	_ = gender         // TODO: delete
	_ = email          // TODO: delete

	// put birthday in YYYY-MM-DD format before storing in db
	var bday string
	for i, s := range strings.Split(birthday, ".") {
		if i == 0 {
			bday = s + bday
		} else {
			bday = s + "-" + bday
		}
	}

	// TODO: register player to the database once repo based code is implemented
}

// checkName takes in a proposed username and returns an integer representing the approval status of the given name.
func checkName(username string) int {
	switch {
	// TODO: come back and re-implement this once repo code is implemented
	// case p.Repo.PlayerExists(username):
	//	return NameAlreadyReserved
	case len(username) > 16:
		return NameTooLong
	case len(username) < 1:
		return NameTooShort
	case !text.ContainsAllowedChars(strings.ToLower(username), text.AllowedCharacters) || strings.Contains(username, " "):
		return NameUnacceptable
	case strings.Contains(strings.ToUpper(username), "MOD-"):
		return NameUnacceptable
	default:
		return NameOk
	}
}

// checkPassword takes in a proposed password and returns an integer representing the approval status of the given password.
func checkPassword(username, password string) int {
	switch {
	case len(password) < 6:
		return PasswordTooShort // too short
	case len(password) > 16:
		return PasswordTooLong // too long
	case !text.ContainsAllowedChars(strings.ToLower(password), text.AllowedCharacters) || strings.Contains(username, " "):
		return PasswordUnacceptable // using non-permitted characters
	case !text.ContainsNumber(password):
		return PasswordHasNoNum // password does not contain a number
	case strings.Contains(password, username):
		return PasswordSimilarToName // name and pass too similar
	default:
		return PasswordOk
	}
}
