package requests

import (
	"github.com/jtieri/habbgo/models"
	"github.com/jtieri/habbgo/services"
)

type CategoryMaxVisitorsReq struct {
	CategoryID int
	Response   chan int
}

func (r *CategoryMaxVisitorsReq) Respond(s services.Servicer) {
	rs := s.(*services.RoomService)
	maxVisitors := rs.MaxVisitorsForCategory(r.CategoryID)
	r.Response <- maxVisitors
}

type CategoryCurrentVisitorsReq struct {
	CategoryID int
	Response   chan int
}

func (r *CategoryCurrentVisitorsReq) Respond(s services.Servicer) {
	rs := s.(*services.RoomService)
	currentVisitors := rs.CurrentVisitorsForCategory(r.CategoryID)
	r.Response <- currentVisitors
}

type PublicRoomsReq struct {
	Response chan []models.Room
}

func (r *PublicRoomsReq) Respond(s services.Servicer) {
	rs := s.(*services.RoomService)
	publicRooms := rs.PublicRooms()
	r.Response <- publicRooms
}
