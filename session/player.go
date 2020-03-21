package session

const (
	RankRegular Rank = "regular"
	RankAdmin        = "admin"
)

type Rank string

type Player struct {
	Name      string
	PrivateID string
	PublicID  string
	Rank      Rank
}
