package models

import "time"

type Player struct {
	ID           int
	Username     string
	Figure       string
	Sex          string // TODO: this should probably be an enum since there are finite values for the string.
	Motto        string
	ConsoleMotto string
	Tickets      int
	PoolFigure   string
	Film         int
	Credits      int
	LastOnline   time.Time
	PlayerRank   Rank
	Badges       []string
	CurrentBadge string
	DisplayBadge bool
	SoundEnabled bool
	LoggedIn     bool
}

type PlayerRoomState struct{}
