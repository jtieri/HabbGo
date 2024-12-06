package protocol

// Incoming handshake packet headers for v14.
var (
	InitCrytpo           = PacketHeader{Name: "INIT_CRYPTO", HeaderID: 206}
	GenerateKey          = PacketHeader{Name: "GENERATEKEY", HeaderID: 202}
	GenerateKeyNew       = PacketHeader{Name: "GENERATEKEY", HeaderID: 2002}
	GetSessionParameters = PacketHeader{Name: "GET_SESSION_PARAMETERS", HeaderID: 181}
	VersionCheck         = PacketHeader{Name: "VERSIONCHECK", HeaderID: 5}
	UniqueID             = PacketHeader{Name: "UNIQUEID", HeaderID: 6}
	SecretKeyIncoming    = PacketHeader{Name: "SECRETKEY", HeaderID: 207}
	SSO                  = PacketHeader{Name: "SSO", HeaderID: 204}
	TryLogin             = PacketHeader{Name: "TRY_LOGIN", HeaderID: 4}
)

// Incoming registration packet headers for v14.
var (
	GetAvailableSets = PacketHeader{Name: "GETAVAILABLESETS", HeaderID: 9}
	GDate            = PacketHeader{Name: "GDATE", HeaderID: 49}
	ApproveName      = PacketHeader{Name: "APPROVENAME", HeaderID: 42}
	ApprovePassword  = PacketHeader{Name: "APPROVE_PASSWORD", HeaderID: 203}
	ApproveEmail     = PacketHeader{Name: "APPROVEEMAIL", HeaderID: 197}
	Register         = PacketHeader{Name: "REGISTER", HeaderID: 43}
)

// Incoming player packet headers for v14.
var (
	GetInfo            = PacketHeader{Name: "GET_INFO", HeaderID: 7}
	GetCredits         = PacketHeader{Name: "GET_CREDITS", HeaderID: 8}
	GetAvailableBadges = PacketHeader{Name: "GETAVAILABLEBADGES", HeaderID: 157}
	GetSoundSetting    = PacketHeader{Name: "GET_SOUND_SETTING", HeaderID: 228}
	TestLatency        = PacketHeader{Name: "TEST_LATENCY", HeaderID: 315}
)

// Incoming navigator packet headers for v14.
var (
	Navigate          = PacketHeader{Name: "NAVIGATE", HeaderID: 150}
	GetUserFlatCats   = PacketHeader{Name: "GETUSERFLATCATS", HeaderID: 151}
	GetFlatInfo       = PacketHeader{Name: "GETFLATINFO", HeaderID: 21}
	Sbuysf            = PacketHeader{Name: "SBUSYF", HeaderID: 13}
	Suserf            = PacketHeader{Name: "SUSERF", HeaderID: 16}
	Srchf             = PacketHeader{Name: "SRCHF", HeaderID: 17}
	Getfvrf           = PacketHeader{Name: "GETFVRF", HeaderID: 18}
	AddFavoriteRoom   = PacketHeader{Name: "ADD_FAVORITE_ROOM", HeaderID: 19}
	DelFavoriteRoom   = PacketHeader{Name: "DEL_FAVORITE_ROOM", HeaderID: 20}
	DeleteFlat        = PacketHeader{Name: "DELETEFLAT", HeaderID: 23}
	UpdateFlat        = PacketHeader{Name: "UPDATEFLAT", HeaderID: 24}
	SetFlatInfo       = PacketHeader{Name: "SETFLATINFO", HeaderID: 25}
	GetFlatCat        = PacketHeader{Name: "GETFLATCAT", HeaderID: 152}
	SetFlatCat        = PacketHeader{Name: "SETFLATCAT", HeaderID: 153}
	GetSpaceNodeUsers = PacketHeader{Name: "GETSPACENODEUSERS", HeaderID: 154}
	RemoveAllRights   = PacketHeader{Name: "REMOVEALLRIGHTS", HeaderID: 155}
	GetParentChain    = PacketHeader{Name: "GETPARENTCHAIN", HeaderID: 156}

	/*
	   tCmds.setaProp("SBUSYF", 13)
	   tCmds.setaProp("SUSERF", 16)
	   tCmds.setaProp("SRCHF", 17)
	   tCmds.setaProp("GETFVRF", 18)
	   tCmds.setaProp("ADD_FAVORITE_ROOM", 19)
	   tCmds.setaProp("DEL_FAVORITE_ROOM", 20)
	   tCmds.setaProp("GETFLATINFO", 21)
	   tCmds.setaProp("DELETEFLAT", 23)
	   tCmds.setaProp("UPDATEFLAT", 24)
	   tCmds.setaProp("SETFLATINFO", 25)
	   tCmds.setaProp("NAVIGATE", 150)
	   tCmds.setaProp("GETUSERFLATCATS", 151)
	   tCmds.setaProp("GETFLATCAT", 152)
	   tCmds.setaProp("SETFLATCAT", 153)
	   tCmds.setaProp("GETSPACENODEUSERS", 154)
	   tCmds.setaProp("REMOVEALLRIGHTS", 155)
	   tCmds.setaProp("GETPARENTCHAIN", 156)
	*/
)

// Incoming recycler packet headers for v14.
var (
	GetFurniRecyclerConfiguration = PacketHeader{Name: "GET_FURNI_RECYCLER_CONFIGURATION", HeaderID: 222}
	GetFurniRecyclerStatus        = PacketHeader{Name: "GET_FURNI_RECYCLER_STATUS", HeaderID: 223}
	ApproveRecycledFurni          = PacketHeader{Name: "APPROVE_RECYCLED_FURNI", HeaderID: 224}
	StartFurniRecycler            = PacketHeader{Name: "START_FURNI_RECYCLING", HeaderID: 225}
	ConfirmFurniRecycling         = PacketHeader{Name: "CONFIRM_FURNI_RECYCLING", HeaderID: 226}

	/*
	   tCmds.setaProp("GET_FURNI_RECYCLER_CONFIGURATION", 222)
	   tCmds.setaProp("GET_FURNI_RECYCLER_STATUS", 223)
	   tCmds.setaProp("APPROVE_RECYCLED_FURNI", 224)
	   tCmds.setaProp("START_FURNI_RECYCLING", 225)
	   tCmds.setaProp("CONFIRM_FURNI_RECYCLING", 226)
	*/
)

// Incoming messenger packet headers for v14.
var (
	MessengerInit          = PacketHeader{Name: "MESSENGERINIT", HeaderID: 12}
	MessengerUpdate        = PacketHeader{Name: "MESSENGER_UPDATE", HeaderID: 15}
	MessengerCClick        = PacketHeader{Name: "MESSENGER_C_CLICK", HeaderID: 30}
	MessengerCRead         = PacketHeader{Name: "MESSENGER_C_READ", HeaderID: 31}
	MessengerMarkRead      = PacketHeader{Name: "MESSENGER_MARKREAD", HeaderID: 32}
	MessengerSendMsg       = PacketHeader{Name: "MESSENGER_SENDMSG", HeaderID: 33}
	MessengerAssignPersMsg = PacketHeader{Name: "MESSENGER_ASSIGNPERSMSG", HeaderID: 36}
	MessengerAcceptBuddy   = PacketHeader{Name: "MESSENGER_ACCEPTBUDDY", HeaderID: 37}
	MessengerDeclineBuddy  = PacketHeader{Name: "MESSENGER_DECLINEBUDDY", HeaderID: 38}
	MessengerRequestBuddy  = PacketHeader{Name: "MESSENGER_REQUESTBUDDY", HeaderID: 39}
	MessengerRemoveBuddy   = PacketHeader{Name: "MESSENGER_REMOVEBUDDY", HeaderID: 40}
	FindUser               = PacketHeader{Name: "FINDUSER", HeaderID: 41}
	MessengerGetMessage    = PacketHeader{Name: "MESSENGER_GETMESSAGES", HeaderID: 191}
	MessengerReportMessage = PacketHeader{Name: "MESSENGER_REPORTMESSAGE", HeaderID: 201}
	/*
	  tCmds.setaProp("MESSENGERINIT", 12)
	  tCmds.setaProp("MESSENGER_UPDATE", 15)
	  tCmds.setaProp("MESSENGER_C_CLICK", 30)
	  tCmds.setaProp("MESSENGER_C_READ", 31)
	  tCmds.setaProp("MESSENGER_MARKREAD", 32)
	  tCmds.setaProp("MESSENGER_SENDMSG", 33)
	  tCmds.setaProp("MESSENGER_ASSIGNPERSMSG", 36)
	  tCmds.setaProp("MESSENGER_ACCEPTBUDDY", 37)
	  tCmds.setaProp("MESSENGER_DECLINEBUDDY", 38)
	  tCmds.setaProp("MESSENGER_REQUESTBUDDY", 39)
	  tCmds.setaProp("MESSENGER_REMOVEBUDDY", 40)
	  tCmds.setaProp("FINDUSER", 41)
	  tCmds.setaProp("MESSENGER_GETMESSAGES", 191)
	  tCmds.setaProp("MESSENGER_REPORTMESSAGE", 201)
	*/
)

// Incoming Habbo Club packet headers for v14.
var (
	ScrGetUserInfo  = PacketHeader{Name: "SCR_GET_USER_INFO", HeaderID: 26}
	ScrBuy          = PacketHeader{Name: "SCR_BUY", HeaderID: 190}
	ScrGiftApproval = PacketHeader{Name: "SCR_GIFT_APPROVAL", HeaderID: 210}
	/*
	  tCmds.setaProp("SCR_GET_USER_INFO", 26)
	  tCmds.setaProp("SCR_BUY", 190)
	  tCmds.setaProp("SCR_GIFT_APPROVAL", 210)
	*/
)

// Incoming room packet headers for v14.
var (
	RoomDirectory         = PacketHeader{Name: "ROOM_DIRECTORY", HeaderID: 2}
	GetDoorFlat           = PacketHeader{Name: "GETDOORFLAT", HeaderID: 28}
	Chat                  = PacketHeader{Name: "CHAT", HeaderID: 52}
	Shout                 = PacketHeader{Name: "SHOUT", HeaderID: 55}
	Whisper               = PacketHeader{Name: "WHISPER", HeaderID: 56}
	Quit                  = PacketHeader{Name: "QUIT", HeaderID: 53}
	GoViaDoor             = PacketHeader{Name: "GOVIADOOR", HeaderID: 54}
	TryFlat               = PacketHeader{Name: "TRYFLAT", HeaderID: 57}
	GoToFlat              = PacketHeader{Name: "GOTOFLAT", HeaderID: 59}
	GHmap                 = PacketHeader{Name: "G_HMAP", HeaderID: 60}
	GUsrs                 = PacketHeader{Name: "G_USRS", HeaderID: 61}
	GObjs                 = PacketHeader{Name: "G_OBJS", HeaderID: 62}
	GItems                = PacketHeader{Name: "G_ITEMS", HeaderID: 63}
	GStat                 = PacketHeader{Name: "G_STAT", HeaderID: 64}
	GetStrip              = PacketHeader{Name: "GETSTRIP", HeaderID: 65}
	FlatPropByItem        = PacketHeader{Name: "FLATPROPBYITEM", HeaderID: 66}
	AddStripItem          = PacketHeader{Name: "ADDSTRIPITEM", HeaderID: 67}
	TradeUnaccept         = PacketHeader{Name: "TRADE_UNACCEPT", HeaderID: 68}
	TradeAcceptIncoming   = PacketHeader{Name: "TRADE_ACCEPT", HeaderID: 69}
	TradeCloseIncoming    = PacketHeader{Name: "TRADE_CLOSE", HeaderID: 70}
	TradeOpen             = PacketHeader{Name: "TRADE_OPEN", HeaderID: 71}
	TradeAddItem          = PacketHeader{Name: "TRADE_ADDITEM", HeaderID: 72}
	MoveStuff             = PacketHeader{Name: "MOVESTUFF", HeaderID: 73}
	SetStuffData          = PacketHeader{Name: "SETSTUFFDATA", HeaderID: 74}
	Move                  = PacketHeader{Name: "MOVE", HeaderID: 75}
	ThrowDice             = PacketHeader{Name: "THROW_DICE", HeaderID: 76}
	DiceOff               = PacketHeader{Name: "DICE_OFF", HeaderID: 77}
	PresentOpenIncoming   = PacketHeader{Name: "PRESENTOPEN", HeaderID: 78}
	LookTo                = PacketHeader{Name: "LOOKTO", HeaderID: 79}
	CarryDrink            = PacketHeader{Name: "CARRYDRINK", HeaderID: 80}
	IntoDoor              = PacketHeader{Name: "INTODOOR", HeaderID: 81}
	DoorGoIn              = PacketHeader{Name: "DOORGOIN", HeaderID: 82}
	GIdata                = PacketHeader{Name: "G_IDATA", HeaderID: 83}
	SetItemData           = PacketHeader{Name: "SETITEMDATA", HeaderID: 84}
	RemoveItemIncoming    = PacketHeader{Name: "REMOVEITEM", HeaderID: 85}
	CarryItem             = PacketHeader{Name: "CARRYITEM", HeaderID: 87}
	Stop                  = PacketHeader{Name: "STOP", HeaderID: 88}
	UseItem               = PacketHeader{Name: "USEITEM", HeaderID: 89}
	PlaceStuff            = PacketHeader{Name: "PLACESTUFF", HeaderID: 90}
	Dance                 = PacketHeader{Name: "DANCE", HeaderID: 93}
	Wave                  = PacketHeader{Name: "WAVE", HeaderID: 94}
	KickUser              = PacketHeader{Name: "KICKUSER", HeaderID: 95}
	AssignRights          = PacketHeader{Name: "ASSIGNRIGHTS", HeaderID: 96}
	RemoveRights          = PacketHeader{Name: "REMOVERIGHTS", HeaderID: 97}
	LetUserIn             = PacketHeader{Name: "LETUSERIN", HeaderID: 98}
	RemoveStuff           = PacketHeader{Name: "REMOVESTUFF", HeaderID: 99}
	GoAway                = PacketHeader{Name: "GOAWAY", HeaderID: 115}
	GetRoomAd             = PacketHeader{Name: "GETROOMAD", HeaderID: 126}
	GetPetStat            = PacketHeader{Name: "GETPETSTAT", HeaderID: 128}
	SetBadge              = PacketHeader{Name: "SETBADGE", HeaderID: 158}
	GetInterst            = PacketHeader{Name: "GETINTERST", HeaderID: 182}
	ConvertFurniToCredits = PacketHeader{Name: "CONVERT_FURNI_TO_CREDITS", HeaderID: 183}
	RoomQueueChange       = PacketHeader{Name: "ROOM_QUEUE_CHANGE", HeaderID: 211}
	SetItemState          = PacketHeader{Name: "SETITEMSTATE", HeaderID: 214}
	GetSpectatorAmount    = PacketHeader{Name: "GET_SPECTATOR_AMOUNT", HeaderID: 216}
	GetGroupBadges        = PacketHeader{Name: "GET_GROUP_BADGES", HeaderID: 230}
	GetGroupDetails       = PacketHeader{Name: "GET_GROUP_DETAILS", HeaderID: 231}
	SpinWheelOfFortune    = PacketHeader{Name: "SPIN_WHEEL_OF_FORTUNE", HeaderID: 247}
	RateFlat              = PacketHeader{Name: "RATEFLAT", HeaderID: 261}
)

/*
  tCmds.setaProp(#room_directory, 2)
  tCmds.setaProp("GETDOORFLAT", 28)
  tCmds.setaProp("CHAT", 52)
  tCmds.setaProp("SHOUT", 55)
  tCmds.setaProp("WHISPER", 56)
  tCmds.setaProp("QUIT", 53)
  tCmds.setaProp("GOVIADOOR", 54)
  tCmds.setaProp("TRYFLAT", 57)
  tCmds.setaProp("GOTOFLAT", 59)
  tCmds.setaProp("G_HMAP", 60)
  tCmds.setaProp("G_USRS", 61)
  tCmds.setaProp("G_OBJS", 62)
  tCmds.setaProp("G_ITEMS", 63)
  tCmds.setaProp("G_STAT", 64)
  tCmds.setaProp("GETSTRIP", 65)
  tCmds.setaProp("FLATPROPBYITEM", 66)
  tCmds.setaProp("ADDSTRIPITEM", 67)
  tCmds.setaProp("TRADE_UNACCEPT", 68)
  tCmds.setaProp("TRADE_ACCEPT", 69)
  tCmds.setaProp("TRADE_CLOSE", 70)
  tCmds.setaProp("TRADE_OPEN", 71)
  tCmds.setaProp("TRADE_ADDITEM", 72)
  tCmds.setaProp("MOVESTUFF", 73)
  tCmds.setaProp("SETSTUFFDATA", 74)
  tCmds.setaProp("MOVE", 75)
  tCmds.setaProp("THROW_DICE", 76)
  tCmds.setaProp("DICE_OFF", 77)
  tCmds.setaProp("PRESENTOPEN", 78)
  tCmds.setaProp("LOOKTO", 79)
  tCmds.setaProp("CARRYDRINK", 80)
  tCmds.setaProp("INTODOOR", 81)
  tCmds.setaProp("DOORGOIN", 82)
  tCmds.setaProp("G_IDATA", 83)
  tCmds.setaProp("SETITEMDATA", 84)
  tCmds.setaProp("REMOVEITEM", 85)
  tCmds.setaProp("CARRYITEM", 87)
  tCmds.setaProp("STOP", 88)
  tCmds.setaProp("USEITEM", 89)
  tCmds.setaProp("PLACESTUFF", 90)
  tCmds.setaProp("DANCE", 93)
  tCmds.setaProp("WAVE", 94)
  tCmds.setaProp("KICKUSER", 95)
  tCmds.setaProp("ASSIGNRIGHTS", 96)
  tCmds.setaProp("REMOVERIGHTS", 97)
  tCmds.setaProp("LETUSERIN", 98)
  tCmds.setaProp("REMOVESTUFF", 99)
  tCmds.setaProp("GOAWAY", 115)
  tCmds.setaProp("GETROOMAD", 126)
  tCmds.setaProp("GETPETSTAT", 128)
  tCmds.setaProp("SETBADGE", 158)
  tCmds.setaProp("GETINTERST", 182)
  tCmds.setaProp("CONVERT_FURNI_TO_CREDITS", 183)
  tCmds.setaProp("ROOM_QUEUE_CHANGE", 211)
  tCmds.setaProp("SETITEMSTATE", 214)
  tCmds.setaProp("GET_SPECTATOR_AMOUNT", 216)
  tCmds.setaProp("GET_GROUP_BADGES", 230)
  tCmds.setaProp("GET_GROUP_DETAILS", 231)
  tCmds.setaProp("SPIN_WHEEL_OF_FORTUNE", 247)
  tCmds.setaProp("RATEFLAT", 261)
*/
