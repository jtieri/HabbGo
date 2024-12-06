package outgoing

import (
	"strconv"

	"github.com/jtieri/habbgo/protocol"
)

// TODO: these should probably be an enum since there are a finite number of values the client expects.
const ( // Used in SESSIONPARAMETERS
	registerCoppa              = 0 // toggle conf_coppa or conf_strong_coppa_req by setting value > 0 or > 1
	voucherEnabled             = 1 // Enables in-game vouchers when value is set > 0
	registerRequireParentEmail = 2 // Requires parent email when registering if value is set > 0
	registerSendParentEmail    = 3 // conf_parent_email_request_reregistration
	allowDirectMail            = 4 // conf_allow_direct_mail
	dateFormat                 = 5 // Sets the date formatter used across the client
	partnerIntegrationEnabled  = 6 // conf_partner_integration. Value is either 1 or 0 (enabled or disabled)
	allowProfileEditing        = 7 // Enables the in-game profile editor
	trackingHeader             = 8 // tracking_header - used in stats.tracking.javascript(?)
	tutorialEnabled            = 9 // Enables the in-game tutorial when value is set to 1 and disables it when 0
)

const (
	// TODO: make this a configurable option.
	availableSets = "[100,105,110,115,120,125,130,135,140,145,150,155,160,165,170,175,176,177,178,180,185,190,195,200,205,206,207,210,215,220,225,230,235,240,245,250,255,260,265,266,267,270,275,280,281,285,290,295,300,305,500,505,510,515,520,525,530,535,540,545,550,555,565,570,575,580,585,590,595,596,600,605,610,615,620,625,626,627,630,635,640,645,650,655,660,665,667,669,670,675,680,685,690,695,696,700,705,710,715,720,725,730,735,740]"
)

// Base64 Header @@
func Hello() protocol.OutgoingPacket {
	return protocol.NewOutgoing(protocol.Hello)
}

// Base64 Header DU
func CryptoParameters() protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.CryptoParameters)
	packet.WriteInt(0) // Toggles server->client encryption; 0=off & non-zero=on
	return packet
}

// TODO: fix this function because currently it does nothing.
func SecretKey() protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.SecretKeyOutgoing)
	packet.WriteString("dsfsfaefsadfdsffdshdsfgfdfdsafdasefasdfasdfsdgfdsgdsfgsdfgds")
	return packet
}

// Base 64 Header DV
func EndCrypto() protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.EndCrypto)
	return packet
}

// Base64 Header DA
func SessionParameters() protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.SessionParameters)

	// TODO: make all of these values configurable instead of using hard-coded values.
	params := make(map[int]string, 9)
	params[voucherEnabled] = strconv.Itoa(0)
	params[registerRequireParentEmail] = strconv.Itoa(0)
	params[registerSendParentEmail] = strconv.Itoa(0)
	params[allowDirectMail] = strconv.Itoa(0)
	params[dateFormat] = "dd-MM-yyyy"
	params[partnerIntegrationEnabled] = strconv.Itoa(0)
	params[allowProfileEditing] = strconv.Itoa(1)
	params[trackingHeader] = ""
	params[tutorialEnabled] = strconv.Itoa(0) // TODO check if player has finished tutorial then set appropriately

	packet.WriteInt(len(params))

	for i, v := range params {
		packet.WriteInt(i)

		if isNumber(v) {
			num, _ := strconv.Atoi(v)
			packet.WriteInt(num)
		} else {
			packet.WriteString(v)
		}
	}

	return packet
}

// Base64 Header "@H"
func AvailableSets() protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.AvailableSets)
	packet.Write(availableSets)
	return packet
}

// Base64 Header @C
func LoginOK() protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.LoginOk)
	return packet
}

// Base64 Header @a
func LocalisedError(errorMsg string) protocol.OutgoingPacket {
	packet := protocol.NewOutgoing(protocol.LocalisedError)
	packet.Write(errorMsg)
	return packet
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}
