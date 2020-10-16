package session

const (
	PrivelegeRegular Privelege = 0
	PrivelegeAdmin             = 1
)

type PlayerID string

type Privelege int

type Player struct {
	ID   PlayerID
	Name string
}
