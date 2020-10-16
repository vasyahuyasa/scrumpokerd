package session

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type SessionID string

type Voice struct {
	PlayerID PlayerID
	Points   float32
}

type Vote struct {
	Title   string
	StartAt time.Time
	Voices  []Voice
}

type PrivelegedPlayer struct {
	Player
	Privelege Privelege
}

type Session struct {
	ID          SessionID
	Title       string
	Players     []PrivelegedPlayer
	CurrentVote Vote
}

func (s *Session) addPlayer(player Player, priv Privelege) error {
	for _, p := range s.Players {
		if p.ID == player.ID {
			return errors.New("player with same id already in session")
		}
	}

	s.Players = append(s.Players, PrivelegedPlayer{
		Player:    player,
		Privelege: priv,
	})

	return nil
}

func generateSessionID() SessionID {
	id := uuid.NewV4()

	return SessionID(id.String())
}
