package incoming

import (
	"fmt"

	"github.com/jtieri/habbgo/models"
	"github.com/jtieri/habbgo/protocol"
	"github.com/jtieri/habbgo/protocol/outgoing"
	"github.com/jtieri/habbgo/services"
	"github.com/jtieri/habbgo/services/requests"
)

func Navigate(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	hideFullRooms := packet.ReadInt() == 1
	categoryID := packet.ReadInt()

	category := getCategory(categoryID, queues.NavigatorQueues)
	subCategories := getSubCategories(categoryID, queues.NavigatorQueues)
	currentVisitors := getCurrentVisitors(categoryID, queues.RoomQueues)
	maxVisitors := getMaxVisitors(categoryID, queues.RoomQueues)
	publicRooms := getPublicRooms(queues.RoomQueues)

	var rooms []models.Room
	if category.IsPublic {
		for _, room := range publicRooms {
			if room.Hidden {
				continue
			}

			if room.CategoryID != category.ID {
				continue
			}

			if hideFullRooms && room.CurrentVisitors >= room.MaxVisitors {
				continue
			}

			rooms = append(rooms, room)
		}
	}

	fmt.Printf("Number of public rooms: %d \n", len(publicRooms))
	fmt.Printf("Number of rooms: %d \n", len(rooms))

	// TODO: finish logic for private rooms

	subCategoryInfo := make([]outgoing.SubCategoryInfo, len(subCategories))

	for i, subCategory := range subCategories {
		currentVisitors := getCurrentVisitors(subCategory.ID, queues.RoomQueues)
		maxVisitors := getMaxVisitors(subCategory.ID, queues.RoomQueues)

		subCategoryInfo[i] = outgoing.SubCategoryInfo{
			ID:              subCategory.ID,
			Name:            subCategory.Name,
			CurrentVisitors: currentVisitors,
			MaxVisitors:     maxVisitors,
			MinRankAccess:   subCategory.MinRankAccess,
		}
	}

	// TODO sort rooms by player count before sending NAVNODEINFO

	session.Send(outgoing.NavNodeInfo(category, subCategoryInfo, rooms, hideFullRooms, currentVisitors, maxVisitors, "test", 1))
}

// GetUserFlatCats is sent from the client requesting the Navigator private room categories that
// should be visible for the specified user.
func GetUserFlatCats(session Session, packet protocol.IncomingPacket, queues *services.ServiceQueues) {
	rankReq := &requests.PlayerRankRequest{
		SessionID: session.ID(),
		Response:  make(chan models.Rank),
	}

	queues.PlayerQueues.QueueRequest(rankReq)
	rank := <-rankReq.Response

	catReq := &requests.PrivateCategoriesForUserReq{
		SessionID:  session.ID(),
		PlayerRank: rank,
		Response:   make(chan []models.Category),
	}

	queues.NavigatorQueues.QueueRequest(catReq)
	cats := <-catReq.Response

	session.Send(outgoing.UserFlatCats(cats))
}

func getCategory(categoryID int, queue services.ServiceQueue) models.Category {
	navCatReq := &requests.NavigatorCategoryReq{
		CategoryID: categoryID,
		Response:   make(chan models.Category),
	}

	queue.QueueRequest(navCatReq)

	category := <-navCatReq.Response

	if category.Name == "" {
		// TODO: category did not exist, handle this case gracefully.
		fmt.Printf("Navigator category not found: %d \n", categoryID)
	}

	return category
}

func getSubCategories(categoryID int, queue services.ServiceQueue) []models.Category {
	subCatReq := &requests.CategoriesByParentIDReq{
		ParentID: categoryID,
		Response: make(chan []models.Category),
	}

	queue.QueueRequest(subCatReq)

	subCategories := <-subCatReq.Response
	if len(subCategories) == 0 {
		// TODO: category did not exist, handle this case gracefully.
		fmt.Printf("Navigator sub-categories not found: %d \n", categoryID)
	}

	return subCategories
}

func getCurrentVisitors(categoryID int, queue services.ServiceQueue) int {
	currentVisitorReq := &requests.CategoryCurrentVisitorsReq{
		CategoryID: categoryID,
		Response:   make(chan int),
	}

	queue.QueueRequest(currentVisitorReq)
	return <-currentVisitorReq.Response
}

func getMaxVisitors(categoryID int, queue services.ServiceQueue) int {
	maxVisitorReq := &requests.CategoryMaxVisitorsReq{
		CategoryID: categoryID,
		Response:   make(chan int),
	}

	queue.QueueRequest(maxVisitorReq)
	return <-maxVisitorReq.Response
}

func getPublicRooms(queue services.ServiceQueue) []models.Room {
	publicRoomReq := &requests.PublicRoomsReq{
		Response: make(chan []models.Room),
	}

	queue.QueueRequest(publicRoomReq)
	return <-publicRoomReq.Response
}
