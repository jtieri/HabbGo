package database

import (
	"database/sql"
	"log"

	"github.com/jtieri/habbgo/models"
)

type PlayerRepo struct {
	database *sql.DB
}

// NewPlayerRepo returns a new instance of PlayerRepo which Player's utilize for accessing the database.
func NewPlayerRepo(db *sql.DB) *PlayerRepo {
	return &PlayerRepo{database: db}
}

// TODO: implement me
func (pr *PlayerRepo) Register(username, figure, gender, email, birthday, createdAt, password string, salt []byte) error {
	return nil
}

// TODO: implement me
func (pr *PlayerRepo) Login(username string, password string) (models.Player, error) {
	return models.Player{}, nil
}

// TODO: implement me
func (pr *PlayerRepo) LoadBadges(playerID int) []string {
	return []string{}
}

// TODO: implement me
func (pr *PlayerRepo) PlayerExists(username string) bool {
	return false
}

func (pr *PlayerRepo) fillDetails(p *models.Player) {
	query := "SELECT P.id, P.username, P.sex, P.figure, P.pool_figure, P.film, P.credits, P.tickets, P.motto, " +
		"P.console_motto, P.last_online, P.sound_enabled, P.Rank " +
		"FROM Players P " +
		"WHERE P.username = $1"

	var tmpRank int
	err := pr.database.QueryRow(query, p.Username).Scan(
		&p.ID,
		&p.Username,
		&p.Sex,
		&p.Figure,
		&p.PoolFigure,
		&p.Film,
		&p.Credits,
		&p.Tickets,
		&p.Motto,
		&p.ConsoleMotto,
		&p.LastOnline,
		&p.SoundEnabled,
		&tmpRank,
	)

	if err != nil {
		log.Printf("%v ", err) // TODO handle & log database errors properly
	}

	p.PlayerRank = models.Rank(tmpRank)
}
