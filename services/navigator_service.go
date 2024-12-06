package services

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/jtieri/habbgo/models"
	"github.com/jtieri/habbgo/repo/database"
)

type NavigatorService struct {
	log *slog.Logger

	queues ServiceQueue
	repo   NavigatorRepository

	errChan      chan error
	shutdownChan chan bool

	navigator models.Navigator
	running   bool
}

// NewNavigatorService instantiates a new NavigatorService and starts it in its own goroutine.
func NewNavigatorService(
	ctx context.Context,
	log *slog.Logger,
	errChan chan error,
	shutdown chan bool,
	db *sql.DB,
) ServiceQueue {
	ns := NavigatorService{
		log:          log,
		queues:       NewServiceQueue(),
		repo:         database.NewNavigatorRepo(db),
		errChan:      errChan,
		shutdownChan: shutdown,
		running:      false,
	}

	go ns.start(ctx)

	return ns.queues
}

// start attempts to listen for and handle incoming Message and Request.
// start should only be called once in its own goroutine.
func (ns *NavigatorService) start(ctx context.Context) {
	if ns.running {
		ns.errChan <- errors.New("the NavigatorService is already running")
		return
	}

	categories, err := ns.repo.Categories()
	if err != nil {
		// TODO: handle database related errors.
		// Depending on the error this may result in a retry, a no-op, or the root ctx being cancelled by the game server.

		ns.log.Error("Failed to get categories from database", "error", err)
	}

	ns.navigator = models.NewNavigator(categories)

	ns.running = true

	defer func() {
		ns.running = false
	}()

	ns.log.Debug("NavigatorService is running")

	for {
		select {
		case msg, open := <-ns.queues.msgQueue:
			if !open {
				ns.log.Debug("NavigatorService msg queue is closed")
				return
			}

			ns.log.Debug("NavigatorService received a message from the queue")

			msg.Handle(ns)
		case req, open := <-ns.queues.reqQueue:
			if !open {
				ns.log.Debug("NavigatorService req queue is closed")
				return
			}

			ns.log.Debug("NavigatorService received a request from the queue")

			req.Respond(ns)
		case <-ctx.Done():
			ns.log.Debug("Context cancelled")
			ns.log.Debug("Shutting down NavigatorService...")

			ns.shutdownChan <- true

			return
		}
	}
}

func (ns *NavigatorService) CategoryByID(id int) (models.Category, bool) {
	return ns.navigator.CategoryByID(id)
}

func (ns *NavigatorService) CategoryByParentID(id int) []models.Category {
	return ns.navigator.CategoriesByParentID(id)
}

func (ns *NavigatorService) PrivateCategoriesForPlayerRank(rank models.Rank) []models.Category {
	var categories []models.Category

	for _, category := range ns.navigator.Categories() {
		if category.IsPublic || rank < category.MinRankAccess {
			continue
		}

		categories = append(categories, category)
	}

	return categories
}
