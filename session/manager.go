package session

import (
	"sync"
)

type Manager struct {
	sessionsMu *sync.Mutex
	sessions   map[SessionID]Session
}

func (m *Manager) CreateSession(admin Player, info SessionInfo) {
	sid := generateSessionID()

	s := Session{
		ID:      sid,
		Info:    info,
		Players: []Player{admin},
	}

	m.sessionsMu.Lock()
	m.sessions[sid] = s
	m.sessionsMu.Unlock()
}
