package models

//go:generate stringer -type=Rank

// Rank represents a Player's rank.
type Rank int

const (
	None Rank = iota
	Normal
	CommunityManager
	Guide
	Hobba
	SuperHobba
	Moderator
	Administrator
)

func Ranks() []Rank {
	return []Rank{
		None,
		Normal,
		CommunityManager,
		Guide,
		Hobba,
		SuperHobba,
		Moderator,
		Administrator,
	}
}
