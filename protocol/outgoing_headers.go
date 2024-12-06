package protocol

type PacketHeader struct {
	Name     string
	HeaderID int
}

// Outgoing handshake packet headers for v14.
var (
	Hello             = PacketHeader{Name: "HELLO", HeaderID: 0}
	CryptoParameters  = PacketHeader{Name: "CRYPTOPARAMETERS", HeaderID: 277}
	SecretKeyOutgoing = PacketHeader{Name: "SECRETKEY", HeaderID: 1}
	EndCrypto         = PacketHeader{Name: "ENDCRYPTO", HeaderID: 278}
	SessionParameters = PacketHeader{Name: "SESSIONPARAMETERS", HeaderID: 257}
	AvailableSets     = PacketHeader{Name: "AVAILABESETS", HeaderID: 8}
	LoginOk           = PacketHeader{Name: "LOGINOK", HeaderID: 3}
	LocalisedError    = PacketHeader{Name: "LOCALISED_ERROR", HeaderID: 33}
)

// Outgoing registration packet headers for v14.
var (
	Date             = PacketHeader{Name: "DATE", HeaderID: 163}
	ApproveNameReply = PacketHeader{Name: "APPROVENAMEREPLY", HeaderID: 36}
	NameUnacceptable = PacketHeader{Name: "NAMEUNACCEPTABLE", HeaderID: 37}
	PasswordApproved = PacketHeader{Name: "PASSWORD_APPROVED", HeaderID: 282}
	EmailApproved    = PacketHeader{Name: "EMAIL_APPROVED", HeaderID: 271}
	EmailRejected    = PacketHeader{Name: "EMAIL_REJECTED", HeaderID: 272}
)

// Outgoing player packet headers for v14
var (
	UserObj         = PacketHeader{Name: "USEROBJ", HeaderID: 5}
	CreditBalance   = PacketHeader{Name: "CREDITBALANCE", HeaderID: 6}
	AvailableBadges = PacketHeader{Name: "AVAILABELBADGES", HeaderID: 229}
	SoundSetting    = PacketHeader{Name: "SOUNDSETTINGS", HeaderID: 308}
	Latency         = PacketHeader{Name: "LATENCY", HeaderID: 354}
)

// Outgoing navigator packet headers for v14.
var (
	FlatResults16         = PacketHeader{Name: "FLAT_RESULTS", HeaderID: 16}
	Error                 = PacketHeader{Name: "ERROR", HeaderID: 33}
	FlatInfo              = PacketHeader{Name: "FLATINFO", HeaderID: 54}
	FlatResults55         = PacketHeader{Name: "FLAT_RESULTS", HeaderID: 55}
	NoFlatsForUser        = PacketHeader{Name: "NOFLATSFORUSER", HeaderID: 57}
	NoFlats               = PacketHeader{Name: "NOFLATS", HeaderID: 58}
	FavouriteRoomRestults = PacketHeader{Name: "FAVOURITEROOMRESULTS", HeaderID: 61}
	FlatPasswordOk        = PacketHeader{Name: "FLATPASSWORD_OK", HeaderID: 130}
	NavNodeInfo           = PacketHeader{Name: "NAVNODEINFO", HeaderID: 220}
	UserFlatCats          = PacketHeader{Name: "USERFLATCATS", HeaderID: 221}
	FlatCat               = PacketHeader{Name: "FLATCAT", HeaderID: 222}
	SpaceNodeUsers        = PacketHeader{Name: "SPACENODEUSERS", HeaderID: 223}
	CantConnect           = PacketHeader{Name: "CANTCONNECT", HeaderID: 224}
	Success               = PacketHeader{Name: "SUCCESS", HeaderID: 225}
	Failure               = PacketHeader{Name: "FAILURE", HeaderID: 226}
	Parentchain           = PacketHeader{Name: "PARENTCHAIN", HeaderID: 227}
	RoomForward           = PacketHeader{Name: "ROOMFORWARD", HeaderID: 286}
	/*
	  tMsgs.setaProp(16, #handle_flat_results)
	  tMsgs.setaProp(33, #handle_error)
	  tMsgs.setaProp(54, #handle_flatinfo)
	  tMsgs.setaProp(55, #handle_flat_results)
	  tMsgs.setaProp(57, #handle_noflatsforuser)
	  tMsgs.setaProp(58, #handle_noflats)
	  tMsgs.setaProp(61, #handle_favouriteroomresults)
	  tMsgs.setaProp(130, #handle_flatpassword_ok)
	  tMsgs.setaProp(220, #handle_navnodeinfo)
	  tMsgs.setaProp(221, #handle_userflatcats)
	  tMsgs.setaProp(222, #handle_flatcat)
	  tMsgs.setaProp(223, #handle_spacenodeusers)
	  tMsgs.setaProp(224, #handle_cantconnect)
	  tMsgs.setaProp(225, #handle_success)
	  tMsgs.setaProp(226, #handle_failure)
	  tMsgs.setaProp(227, #handle_parentchain)
	  tMsgs.setaProp(286, #handle_roomforward)
	*/
)

// Outgoing recycler packet headers for v14.
var (
	RecylcerConfiguration  = PacketHeader{Name: "RECYCLER_CONFIGURATION", HeaderID: 303}
	RecyclerStatus         = PacketHeader{Name: "RECYCLER_STATUS", HeaderID: 304}
	ApproveRecyclingResult = PacketHeader{Name: "APPROVE_RECYCLING_RESULT", HeaderID: 305}
	StartRecyclingResults  = PacketHeader{Name: "START_RECYCLING_RESULT", HeaderID: 306}
	ConfirmRecyclingResult = PacketHeader{Name: "CONFIRM_RECYCLING_RESULT", HeaderID: 307}

	/*
	   tMsgs.setaProp(303, #handle_recycler_configuration)
	   tMsgs.setaProp(304, #handle_recycler_status)
	   tMsgs.setaProp(305, #handle_approve_recycling_result)
	   tMsgs.setaProp(306, #handle_start_recycling_result)
	   tMsgs.setaProp(307, #handle_confirm_recycling_result)
	*/
)

// Outgoing messenger packet headers for v14.
var (
/*
tMsgs.setaProp(13, #handle_console_update)

	tMsgs.setaProp(128, #handle_memberinfo)
	tMsgs.setaProp(132, #handle_buddy_request)
	tMsgs.setaProp(133, #handle_campaign_message)
	tMsgs.setaProp(134, #handle_messenger_messages)
	tMsgs.setaProp(137, #handle_add_buddy)
	tMsgs.setaProp(138, #handle_remove_buddy)
	tMsgs.setaProp(147, #handle_mypersistentmessage)
	tMsgs.setaProp(260, #handle_messenger_error)
	tMsgs.setaProp(263, #handle_buddylist)
*/
)

// Outgoing Habbo Club packet headers for v14.
var (
/*
tMsgs.setaProp(3, #handle_ok)

	tMsgs.setaProp(7, #handle_scr_sinfo)
	tMsgs.setaProp(280, #handle_gift)
*/
)

// TODO: sanity check these values!!
// Outgoing room packet handlers for v14.
var (
	Disconnect = PacketHeader{Name: "DISCONNECT", HeaderID: -1}
	Clc        = PacketHeader{Name: "CLC", HeaderID: 18}
	OpcOk      = PacketHeader{Name: "OPC_OK", HeaderID: 19}
	/*
				CHAT Chat24, // @X
		        SHOUT Chat26 // @Z
		        WHISPER Chat25 // @Y
	*/
	Chat24        = PacketHeader{Name: "CHAT", HeaderID: 24}
	Chat25        = PacketHeader{Name: "CHAT", HeaderID: 25}
	Chat26        = PacketHeader{Name: "CHAT", HeaderID: 26}
	Users         = PacketHeader{Name: "USERS", HeaderID: 28}
	Logout        = PacketHeader{Name: "LOGOUT", HeaderID: 29}
	Objects       = PacketHeader{Name: "OBJECTS", HeaderID: 30}
	Heightmap     = PacketHeader{Name: "HEIGHTMAP", HeaderID: 31}
	ActiveObjects = PacketHeader{Name: "ACTIVEOBJECTS", HeaderID: 32}
	// Error                 = PacketHeader{Name: "ERROR", HeaderID: 33}
	Status                = PacketHeader{Name: "STATUS", HeaderID: 34}
	FlatLetIn             = PacketHeader{Name: "FLAT_LETIN", HeaderID: 41}
	Items45               = PacketHeader{Name: "ITEMS", HeaderID: 45}
	RoomRights42          = PacketHeader{Name: "ROOM_RIGHTS", HeaderID: 42}
	RoomRights43          = PacketHeader{Name: "ROOM_RIGHTS", HeaderID: 43}
	FlatProperty          = PacketHeader{Name: "FLATPROPERTY", HeaderID: 46}
	RoomRights47          = PacketHeader{Name: "ROOM_RIGHTS", HeaderID: 47}
	Idata                 = PacketHeader{Name: "IDATA", HeaderID: 48}
	DoorFlat              = PacketHeader{Name: "DOORFLAT", HeaderID: 62}
	DoorDeleted63         = PacketHeader{Name: "DOORDELETED", HeaderID: 63}
	DoorDeleted64         = PacketHeader{Name: "DOORDELETED", HeaderID: 64}
	RoomReady             = PacketHeader{Name: "ROOM_READY", HeaderID: 69}
	YouAreMod             = PacketHeader{Name: "YOUAREMOD", HeaderID: 70}
	ShowProgram           = PacketHeader{Name: "SHOWPROGRAM", HeaderID: 71}
	NoUserForGift         = PacketHeader{Name: "NO_USER_FOR_GIFT", HeaderID: 76}
	Items83               = PacketHeader{Name: "ITEMS", HeaderID: 83}
	RemoveItem            = PacketHeader{Name: "REMOVEITEM", HeaderID: 84}
	UpdateItem            = PacketHeader{Name: "UPDATEITEM", HeaderID: 85}
	StuffDataUpdate       = PacketHeader{Name: "STUFFDATAUPDATE", HeaderID: 88}
	DoorOut               = PacketHeader{Name: "DOOR_OUT", HeaderID: 89}
	DiceValue             = PacketHeader{Name: "DICE_VALUE", HeaderID: 90}
	DoorbellRinging       = PacketHeader{Name: "DOORBELL_RINGING", HeaderID: 91}
	DoorIn                = PacketHeader{Name: "DOOR_IN", HeaderID: 92}
	ActiveObjectAdd       = PacketHeader{Name: "ACTIVEOBJECT_ADD", HeaderID: 93}
	ActiveObjectRemove    = PacketHeader{Name: "ACTIVEOBJECT_REMOVE", HeaderID: 94}
	ActiveObjectUpdate    = PacketHeader{Name: "ACTIVEOBJECT_UPDATE", HeaderID: 95}
	StripInfo98           = PacketHeader{Name: "STRIPINFO", HeaderID: 98}
	RemoveStripItem       = PacketHeader{Name: "REMOVESTRIPITEM", HeaderID: 99}
	StripUpdated          = PacketHeader{Name: "STRIPUPDATED", HeaderID: 101}
	YouAreNotAllowed      = PacketHeader{Name: "YOUARENOTALLOWED", HeaderID: 102}
	OtherNotAllowed       = PacketHeader{Name: "OTHERNOTALLOWED", HeaderID: 103}
	TradeCompleted105     = PacketHeader{Name: "TRADE_COMPLETED", HeaderID: 105}
	TradeItems            = PacketHeader{Name: "TRADE_ITEMS", HeaderID: 108}
	TradeAccept           = PacketHeader{Name: "TRADE_ACCEPT", HeaderID: 109}
	TradeClose            = PacketHeader{Name: "TRADE_CLOSE", HeaderID: 110}
	TradeCompleted112     = PacketHeader{Name: "TRADE_COMPLETED", HeaderID: 112}
	PresentOpen           = PacketHeader{Name: "PRESENTOPEN", HeaderID: 129}
	FlatNotAllowedToEnter = PacketHeader{Name: "FLATNOTALLOWEDTOENTER", HeaderID: 131}
	StripInfo140          = PacketHeader{Name: "STRIPINFO", HeaderID: 140}
	RoomAd                = PacketHeader{Name: "ROOMAD", HeaderID: 208}
	PetStat               = PacketHeader{Name: "PETSTAT", HeaderID: 210}
	HeightMapUpdate       = PacketHeader{Name: "HEIGHTMAPUPDATE", HeaderID: 219}
	UserBadge             = PacketHeader{Name: "USERBADGE", HeaderID: 228}
	SlideObjectBundle     = PacketHeader{Name: "SLIDEOBJECTBUNDLE", HeaderID: 230}
	InterstitialData      = PacketHeader{Name: "INTERSTITIALDATA", HeaderID: 258}
	RoomQueueData         = PacketHeader{Name: "ROOMQUEUEDATA", HeaderID: 259}
	YouAreSpectator       = PacketHeader{Name: "YOUARESPECTATOR", HeaderID: 254}
	RemoveSpecs           = PacketHeader{Name: "REMOVESPECS", HeaderID: 283}
	FigureChange          = PacketHeader{Name: "FIGURE_CHANGE", HeaderID: 266}
	SpectatorAmount       = PacketHeader{Name: "SPECTATOR_AMOUNT", HeaderID: 298}
	GroupBadges           = PacketHeader{Name: "GROUP_BADGES", HeaderID: 309}
	GroupMembershipUpdate = PacketHeader{Name: "GROUP_MEMBERSHIP_UPDATE", HeaderID: 310}
	GroupDetails          = PacketHeader{Name: "GROUP_DETAILS", HeaderID: 311}
	RoomRating            = PacketHeader{Name: "ROOM_RATING", HeaderID: 345}
)

/*
  tMsgs = [:]
  tMsgs.setaProp(-1, #handle_disconnect)
  tMsgs.setaProp(18, #handle_clc)
  tMsgs.setaProp(19, #handle_opc_ok)
  tMsgs.setaProp(24, #handle_chat)
  tMsgs.setaProp(25, #handle_chat)
  tMsgs.setaProp(26, #handle_chat)
  tMsgs.setaProp(28, #handle_users)
  tMsgs.setaProp(29, #handle_logout)
  tMsgs.setaProp(30, #handle_OBJECTS)
  tMsgs.setaProp(31, #handle_heightmap)
  tMsgs.setaProp(32, #handle_activeobjects)
  tMsgs.setaProp(33, #handle_error)
  tMsgs.setaProp(34, #handle_status)
  tMsgs.setaProp(41, #handle_flat_letin)
  tMsgs.setaProp(45, #handle_items)
  tMsgs.setaProp(42, #handle_room_rights)
  tMsgs.setaProp(43, #handle_room_rights)
  tMsgs.setaProp(46, #handle_flatproperty)
  tMsgs.setaProp(47, #handle_room_rights)
  tMsgs.setaProp(48, #handle_idata)
  tMsgs.setaProp(62, #handle_doorflat)
  tMsgs.setaProp(63, #handle_doordeleted)
  tMsgs.setaProp(64, #handle_doordeleted)
  tMsgs.setaProp(69, #handle_room_ready)
  tMsgs.setaProp(70, #handle_youaremod)
  tMsgs.setaProp(71, #handle_showprogram)
  tMsgs.setaProp(76, #handle_no_user_for_gift)
  tMsgs.setaProp(83, #handle_items)
  tMsgs.setaProp(84, #handle_removeitem)
  tMsgs.setaProp(85, #handle_updateitem)
  tMsgs.setaProp(88, #handle_stuffdataupdate)
  tMsgs.setaProp(89, #handle_door_out)
  tMsgs.setaProp(90, #handle_dice_value)
  tMsgs.setaProp(91, #handle_doorbell_ringing)
  tMsgs.setaProp(92, #handle_door_in)
  tMsgs.setaProp(93, #handle_activeobject_add)
  tMsgs.setaProp(94, #handle_activeobject_remove)
  tMsgs.setaProp(95, #handle_activeobject_update)
  tMsgs.setaProp(98, #handle_stripinfo)
  tMsgs.setaProp(99, #handle_removestripitem)
  tMsgs.setaProp(101, #handle_stripupdated)
  tMsgs.setaProp(102, #handle_youarenotallowed)
  tMsgs.setaProp(103, #handle_othernotallowed)
  tMsgs.setaProp(105, #handle_trade_completed)
  tMsgs.setaProp(108, #handle_trade_items)
  tMsgs.setaProp(109, #handle_trade_accept)
  tMsgs.setaProp(110, #handle_trade_close)
  tMsgs.setaProp(112, #handle_trade_completed)
  tMsgs.setaProp(129, #handle_presentopen)
  tMsgs.setaProp(131, #handle_flatnotallowedtoenter)
  tMsgs.setaProp(140, #handle_stripinfo)
  tMsgs.setaProp(208, #handle_roomad)
  tMsgs.setaProp(210, #handle_petstat)
  tMsgs.setaProp(219, #handle_heightmapupdate)
  tMsgs.setaProp(228, #handle_userbadge)
  tMsgs.setaProp(230, #handle_slideobjectbundle)
  tMsgs.setaProp(258, #handle_interstitialdata)
  tMsgs.setaProp(259, #handle_roomqueuedata)
  tMsgs.setaProp(254, #handle_youarespectator)
  tMsgs.setaProp(283, #handle_removespecs)
  tMsgs.setaProp(266, #handle_figure_change)
  tMsgs.setaProp(298, #handle_spectator_amount)
  tMsgs.setaProp(309, #handle_group_badges)
  tMsgs.setaProp(310, #handle_group_membership_update)
  tMsgs.setaProp(311, #handle_group_details)
  tMsgs.setaProp(345, #handle_room_rating)
*/
