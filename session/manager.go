package session

import (
	"fmt"
	"sync"
)

type Manager struct {
	sessionsMu *sync.Mutex
	sessions   map[SessionID]Session
}

func (m *Manager) CreateSession(id SessionID, title string) error {
	m.sessionsMu.Lock()
	defer m.sessionsMu.Unlock()

	if _, ok := m.sessions[id]; ok {
		return fmt.Errorf("session with same id %s already exist", id)
	}

	s := Session{
		ID:    id,
		Title: title,
	}

	m.sessions[id] = s

	return nil
}
