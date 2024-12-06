package incoming

import (
	"sync"

	"github.com/jtieri/habbgo/protocol"
	"github.com/jtieri/habbgo/services"
)

type CommandHandler = func(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues)

type IncomingRegistry struct {
	registeredCommands map[protocol.PacketHeader]CommandHandler
	mux                sync.Mutex
}

func NewIncomingRegistry() *IncomingRegistry {
	r := &IncomingRegistry{
		registeredCommands: make(map[protocol.PacketHeader]CommandHandler),
		mux:                sync.Mutex{},
	}

	r.RegisterHandshakeCommands()
	r.RegisterRegistrationCommands()
	r.RegisterPlayerCommands()
	r.RegisterNavigatorCommands()
	r.RegisterRecyclerCommands()
	r.RegisterMessengerCommands()
	r.RegisterHabboClubCommands()
	r.RegisterRoomCommands()

	return r
}

func (r *IncomingRegistry) GetPacketHeader(headerID int) (protocol.PacketHeader, bool) {
	//r.mux.Lock()
	//defer r.mux.Unlock()

	for key, _ := range r.registeredCommands {
		if key.HeaderID == headerID {
			return key, true
		}

		continue
	}

	return protocol.PacketHeader{}, false
}

func (r *IncomingRegistry) GetCommand(packetHeader protocol.PacketHeader) (CommandHandler, bool) {
	//r.mux.Lock()
	//defer r.mux.Unlock()

	h, found := r.registeredCommands[packetHeader]
	return h, found
}

func (r *IncomingRegistry) RegisterHandshakeCommands() {
	r.registeredCommands[protocol.InitCrytpo] = InitCrypto
	r.registeredCommands[protocol.GenerateKey] = GenerateKey    // TODO: figure out all client versions that use this one
	r.registeredCommands[protocol.GenerateKeyNew] = GenerateKey // TODO: figure out all client versions that use this one
	r.registeredCommands[protocol.VersionCheck] = VersionCheck  // 1170 - VERSIONCHECK in later clients? v26+? // TODO figure out exact client revisions when these packet headers change
	r.registeredCommands[protocol.UniqueID] = UniqueID
	r.registeredCommands[protocol.GetSessionParameters] = GetSessionParameters
	r.registeredCommands[protocol.SSO] = SSO
	r.registeredCommands[protocol.TryLogin] = TryLogin
	r.registeredCommands[protocol.SecretKeyIncoming] = SecretKey
}

func (r *IncomingRegistry) RegisterRegistrationCommands() {
	r.registeredCommands[protocol.GetAvailableSets] = GetAvailableSets
	r.registeredCommands[protocol.GDate] = GDate
	r.registeredCommands[protocol.ApproveName] = ApproveName
	r.registeredCommands[protocol.ApprovePassword] = ApprovePassword
	r.registeredCommands[protocol.ApproveEmail] = ApproveEmail
	r.registeredCommands[protocol.Register] = Register
}

func (r *IncomingRegistry) RegisterPlayerCommands() {
	r.registeredCommands[protocol.GetInfo] = GetInfo
	r.registeredCommands[protocol.GetCredits] = GetCredits
	r.registeredCommands[protocol.GetAvailableBadges] = GetAvailableBadges
	r.registeredCommands[protocol.GetSoundSetting] = GetSoundSetting
	r.registeredCommands[protocol.TestLatency] = TestLatency
}

func (r *IncomingRegistry) RegisterNavigatorCommands() {
	r.registeredCommands[protocol.Navigate] = Navigate
	r.registeredCommands[protocol.GetUserFlatCats] = GetUserFlatCats
	r.registeredCommands[protocol.GetFlatInfo] = nil
	r.registeredCommands[protocol.Sbuysf] = nil
	r.registeredCommands[protocol.Suserf] = nil
	r.registeredCommands[protocol.Srchf] = nil
	r.registeredCommands[protocol.Getfvrf] = nil
	r.registeredCommands[protocol.AddFavoriteRoom] = nil
	r.registeredCommands[protocol.DelFavoriteRoom] = nil
	r.registeredCommands[protocol.DeleteFlat] = nil
	r.registeredCommands[protocol.UpdateFlat] = nil
	r.registeredCommands[protocol.SetFlatInfo] = nil
	r.registeredCommands[protocol.GetFlatCat] = nil
	r.registeredCommands[protocol.SetFlatCat] = nil
	r.registeredCommands[protocol.GetSpaceNodeUsers] = nil
	r.registeredCommands[protocol.RemoveAllRights] = nil
	r.registeredCommands[protocol.GetParentChain] = nil
}

func (r *IncomingRegistry) RegisterRecyclerCommands() {
	r.registeredCommands[protocol.GetFurniRecyclerConfiguration] = nil
	r.registeredCommands[protocol.GetFurniRecyclerStatus] = nil
	r.registeredCommands[protocol.ApproveRecycledFurni] = nil
	r.registeredCommands[protocol.StartFurniRecycler] = nil
	r.registeredCommands[protocol.ConfirmFurniRecycling] = nil
}

func (r *IncomingRegistry) RegisterMessengerCommands() {
	r.registeredCommands[protocol.MessengerInit] = nil
	r.registeredCommands[protocol.MessengerUpdate] = nil
	r.registeredCommands[protocol.MessengerCClick] = nil
	r.registeredCommands[protocol.MessengerCRead] = nil
	r.registeredCommands[protocol.MessengerMarkRead] = nil
	r.registeredCommands[protocol.MessengerSendMsg] = nil
	r.registeredCommands[protocol.MessengerAssignPersMsg] = nil
	r.registeredCommands[protocol.MessengerAcceptBuddy] = nil
	r.registeredCommands[protocol.MessengerDeclineBuddy] = nil
	r.registeredCommands[protocol.MessengerRequestBuddy] = nil
	r.registeredCommands[protocol.MessengerRemoveBuddy] = nil
	r.registeredCommands[protocol.FindUser] = nil
	r.registeredCommands[protocol.MessengerGetMessage] = nil
	r.registeredCommands[protocol.MessengerReportMessage] = nil
}

func (r *IncomingRegistry) RegisterHabboClubCommands() {
	r.registeredCommands[protocol.ScrGetUserInfo] = nil
	r.registeredCommands[protocol.ScrBuy] = nil
	r.registeredCommands[protocol.ScrGiftApproval] = nil
}

func (r *IncomingRegistry) RegisterRoomCommands() {
	r.registeredCommands[protocol.RoomDirectory] = RoomDirectory
	r.registeredCommands[protocol.GetInterst] = GetInterst
	r.registeredCommands[protocol.GetRoomAd] = GetRoomAd
	r.registeredCommands[protocol.GHmap] = GHmap
	r.registeredCommands[protocol.GUsrs] = GUsrs
	r.registeredCommands[protocol.GObjs] = GObjs
	r.registeredCommands[protocol.GStat] = GStat
	r.registeredCommands[protocol.GoAway] = GoAway
	r.registeredCommands[protocol.Move] = Move
	r.registeredCommands[protocol.Stop] = Stop
	r.registeredCommands[protocol.Quit] = Quit
}
