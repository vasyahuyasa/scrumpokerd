package session

import (
	uuid "github.com/satori/go.uuid"
)

type SessionID string

type SessionInfo struct {
	Title string
}

type Session struct {
	ID      SessionID
	Info    SessionInfo
	Players []Player
}

func generateSessionID() SessionID {
	id := uuid.NewV4()

	return SessionID(id.String())
}
