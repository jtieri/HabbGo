package requests

import (
	"github.com/jtieri/habbgo/models"
	"github.com/jtieri/habbgo/services"
)

type LoginPlayerReq struct {
	SessionID string
	Username  string
	Password  string
	Response  chan bool
}

func (r *LoginPlayerReq) Respond(s services.Servicer) {
	ps := s.(*services.PlayerService)
	ok := ps.LoginPlayer(r.SessionID, r.Username, r.Password)
	r.Response <- ok
}

type PlayerRankRequest struct {
	SessionID string
	Response  chan models.Rank
}

func (r *PlayerRankRequest) Respond(s services.Servicer) {
	ps := s.(*services.PlayerService)
	rank := ps.PlayerRank(r.SessionID)
	r.Response <- rank
}

type CreditBalanceReq struct {
	SessionID string
	Response  chan int
}

func (r *CreditBalanceReq) Respond(s services.Servicer) {
	ps := s.(*services.PlayerService)
	r.Response <- ps.CreditBalance(r.SessionID)
}

type SoundSettingReq struct {
	SessionID string
	Response  chan bool
}

func (r *SoundSettingReq) Respond(s services.Servicer) {
	ps := s.(*services.PlayerService)
	r.Response <- ps.SoundSetting(r.SessionID)
}

type BadgeInfoReq struct {
	SessionID string
	Response  chan services.BadgeDetails
}

func (r *BadgeInfoReq) Respond(s services.Servicer) {
	ps := s.(*services.PlayerService)
	r.Response <- ps.BadgeInfo(r.SessionID)
}

type UserObjReq struct {
	SessionID string
	Response  chan services.UserObjectInfo
}

func (r *UserObjReq) Respond(s services.Servicer) {
	ps := s.(*services.PlayerService)
	r.Response <- ps.UserObjInfo(r.SessionID)
}
