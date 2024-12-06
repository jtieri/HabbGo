package incoming

import (
	"github.com/jtieri/habbgo/protocol"
	"github.com/jtieri/habbgo/protocol/outgoing"
	"github.com/jtieri/habbgo/services"
	"github.com/jtieri/habbgo/services/requests"
)

func GetInfo(session Session, _ protocol.IncomingPacket, queues *services.ServiceQueues) {
	req := &requests.UserObjReq{
		SessionID: session.ID(),
		Response:  make(chan services.UserObjectInfo),
	}

	queues.PlayerQueues.QueueRequest(req)

	info := <-req.Response

	session.Send(outgoing.UserObj(info))
}

func GetCredits(session Session, _ protocol.IncomingPacket, queues *services.ServiceQueues) {
	req := &requests.CreditBalanceReq{
		SessionID: session.ID(),
		Response:  make(chan int),
	}

	queues.PlayerQueues.QueueRequest(req)

	balance := <-req.Response

	session.Send(outgoing.CreditBalance(balance))
}

func GetAvailableBadges(session Session, _ protocol.IncomingPacket, queues *services.ServiceQueues) {
	req := &requests.BadgeInfoReq{
		SessionID: session.ID(),
		Response:  make(chan services.BadgeDetails),
	}

	queues.PlayerQueues.QueueRequest(req)

	info := <-req.Response

	session.Send(outgoing.AvailableBadges(info))
}

func GetSoundSetting(session Session, _ protocol.IncomingPacket, queues *services.ServiceQueues) {
	req := &requests.SoundSettingReq{
		SessionID: session.ID(),
		Response:  make(chan bool),
	}

	queues.PlayerQueues.QueueRequest(req)

	enabled := <-req.Response

	session.Send(outgoing.SoundSetting(enabled))
}

func TestLatency(session Session, packet protocol.IncomingPacket, _ *services.ServiceQueues) {
	l := packet.ReadInt()
	session.Send(outgoing.Latency(l))
}
