package requests

import (
	"fmt"

	"github.com/jtieri/habbgo/models"
	"github.com/jtieri/habbgo/services"
)

type NavigatorCategoryReq struct {
	CategoryID int
	Response   chan models.Category
}

func (r *NavigatorCategoryReq) Respond(s services.Servicer) {
	catService := s.(*services.NavigatorService)

	category, ok := catService.CategoryByID(r.CategoryID)
	if ok {
		fmt.Printf("FOUND CATEGORY: %d \n", category.ID)
		r.Response <- category
		return
	}

	fmt.Printf("NOT FOUND CATEGORY: %d \n", r.CategoryID)
	r.Response <- models.Category{}
}

type CategoriesByParentIDReq struct {
	ParentID int
	Response chan []models.Category
}

func (r *CategoriesByParentIDReq) Respond(s services.Servicer) {
	catService := s.(*services.NavigatorService)

	categories := catService.CategoryByParentID(r.ParentID)
	r.Response <- categories
}

type PrivateCategoriesForUserReq struct {
	SessionID  string
	PlayerRank models.Rank
	Response   chan []models.Category
}

func (r *PrivateCategoriesForUserReq) Respond(s services.Servicer) {
	ns := s.(*services.NavigatorService)

	categories := ns.PrivateCategoriesForPlayerRank(r.PlayerRank)
	r.Response <- categories
}
