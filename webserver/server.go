package webserver

import (
	"github.com/go-chi/chi"
	"github.com/vasyahuyasa/scrumpokerd/session"
)

type Server struct {
	mux *chi.Mux
	mgr *session.Manager
}

func NewServer(mgr *session.Manager) *Server {
	return &Server{
		mux: chi.NewMux(),
		mgr: mgr,
	}
}

/*
	POST /session - create new session (get id)
	GET /session/{session_id} - session info (get title, players, topic, stage, epoch)
	POST /session/{session_id}/players - join session (send name, get id)
	POST /session/{session_id}/vote - send vote
	GET /session/{session_id}/online
*/

func (s *Server) routes() {
	s.mux.Route("/session", func(r chi.Router) {
		r.Post("/", createSessionHandle)
		r.Get("/{session_id}", getSessionInfohandler)
		r.Post("/{session_id}/players", joinSessionHandler)
		r.Post("/{session_id}/vote", voteHandler)
		r.Get("/{session_id}/online", sseStreamHandler)
	})
}
