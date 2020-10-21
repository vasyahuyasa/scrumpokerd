// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi"
)

// CreatePoll defines model for CreatePoll.
type CreatePoll struct {

	// Poll topic
	Topic string `json:"topic"`
}

// DoVote defines model for DoVote.
type DoVote struct {

	// Vote for current topic
	Points float32 `json:"points"`
}

// JoinSession defines model for JoinSession.
type JoinSession struct {

	// Player name
	Name string `json:"name"`
}

// NewSession defines model for NewSession.
type NewSession struct {

	// Session title
	Title string `json:"title"`
}

// Player defines model for Player.
type Player struct {

	// Player id
	Id *string `json:"id,omitempty"`

	// Name of player
	Name *string `json:"name,omitempty"`
}

// Poll defines model for Poll.
type Poll struct {
	StartAt *string  `json:"startAt,omitempty"`
	Title   *string  `json:"title,omitempty"`
	Voices  *[]Voice `json:"voices,omitempty"`
}

// Voice defines model for Voice.
type Voice struct {
	PlayerID *string  `json:"playerID,omitempty"`
	Points   *float32 `json:"points,omitempty"`
}

// ResponseSessionID defines model for ResponseSessionID.
type ResponseSessionID struct {

	// Title of session
	Title string `json:"title"`
}

// ResponseSessionInfo defines model for ResponseSessionInfo.
type ResponseSessionInfo struct {

	// Session ID
	Id *string `json:"id,omitempty"`

	// List of players in room
	Players *[][]Player `json:"players,omitempty"`
	Poll    *Poll       `json:"poll,omitempty"`

	// Session title
	Title *string `json:"title,omitempty"`
}

// CreateSessionJSONBody defines parameters for CreateSession.
type CreateSessionJSONBody NewSession

// JoinSessionJSONBody defines parameters for JoinSession.
type JoinSessionJSONBody JoinSession

// CreatePollJSONBody defines parameters for CreatePoll.
type CreatePollJSONBody CreatePoll

// VoteJSONBody defines parameters for Vote.
type VoteJSONBody DoVote

// CreateSessionRequestBody defines body for CreateSession for application/json ContentType.
type CreateSessionJSONRequestBody CreateSessionJSONBody

// JoinSessionRequestBody defines body for JoinSession for application/json ContentType.
type JoinSessionJSONRequestBody JoinSessionJSONBody

// CreatePollRequestBody defines body for CreatePoll for application/json ContentType.
type CreatePollJSONRequestBody CreatePollJSONBody

// VoteRequestBody defines body for Vote for application/json ContentType.
type VoteJSONRequestBody VoteJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create new session
	// (POST /session)
	CreateSession(w http.ResponseWriter, r *http.Request)
	// Session data
	// (GET /session/{sessionId})
	GetSession(w http.ResponseWriter, r *http.Request, sessionId string)
	// SSE updates
	// (GET /session/{sessionId}/live)
	LiveStatus(w http.ResponseWriter, r *http.Request, sessionId string)
	// Join session
	// (POST /session/{sessionId}/players)
	JoinSession(w http.ResponseWriter, r *http.Request, sessionId string)
	// Stop current poll
	// (DELETE /session/{sessionId}/poll)
	StopPoll(w http.ResponseWriter, r *http.Request, sessionId string)
	// Create new poll
	// (POST /session/{sessionId}/poll)
	CreatePoll(w http.ResponseWriter, r *http.Request, sessionId string)
	// Vote in current poll
	// (POST /session/{sessionId}/poll/vote)
	Vote(w http.ResponseWriter, r *http.Request, sessionId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateSession operation middleware
func (siw *ServerInterfaceWrapper) CreateSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, "userToken.Scopes", []string{""})

	siw.Handler.CreateSession(w, r.WithContext(ctx))
}

// GetSession operation middleware
func (siw *ServerInterfaceWrapper) GetSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "sessionId" -------------
	var sessionId string

	err = runtime.BindStyledParameter("simple", false, "sessionId", chi.URLParam(r, "sessionId"), &sessionId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sessionId: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, "userToken.Scopes", []string{""})

	siw.Handler.GetSession(w, r.WithContext(ctx), sessionId)
}

// LiveStatus operation middleware
func (siw *ServerInterfaceWrapper) LiveStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "sessionId" -------------
	var sessionId string

	err = runtime.BindStyledParameter("simple", false, "sessionId", chi.URLParam(r, "sessionId"), &sessionId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sessionId: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.LiveStatus(w, r.WithContext(ctx), sessionId)
}

// JoinSession operation middleware
func (siw *ServerInterfaceWrapper) JoinSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "sessionId" -------------
	var sessionId string

	err = runtime.BindStyledParameter("simple", false, "sessionId", chi.URLParam(r, "sessionId"), &sessionId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sessionId: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, "userToken.Scopes", []string{""})

	siw.Handler.JoinSession(w, r.WithContext(ctx), sessionId)
}

// StopPoll operation middleware
func (siw *ServerInterfaceWrapper) StopPoll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "sessionId" -------------
	var sessionId string

	err = runtime.BindStyledParameter("simple", false, "sessionId", chi.URLParam(r, "sessionId"), &sessionId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sessionId: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, "userToken.Scopes", []string{""})

	siw.Handler.StopPoll(w, r.WithContext(ctx), sessionId)
}

// CreatePoll operation middleware
func (siw *ServerInterfaceWrapper) CreatePoll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "sessionId" -------------
	var sessionId string

	err = runtime.BindStyledParameter("simple", false, "sessionId", chi.URLParam(r, "sessionId"), &sessionId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sessionId: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, "userToken.Scopes", []string{""})

	siw.Handler.CreatePoll(w, r.WithContext(ctx), sessionId)
}

// Vote operation middleware
func (siw *ServerInterfaceWrapper) Vote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "sessionId" -------------
	var sessionId string

	err = runtime.BindStyledParameter("simple", false, "sessionId", chi.URLParam(r, "sessionId"), &sessionId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sessionId: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, "userToken.Scopes", []string{""})

	siw.Handler.Vote(w, r.WithContext(ctx), sessionId)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	r.Group(func(r chi.Router) {
		r.Post("/session", wrapper.CreateSession)
	})
	r.Group(func(r chi.Router) {
		r.Get("/session/{sessionId}", wrapper.GetSession)
	})
	r.Group(func(r chi.Router) {
		r.Get("/session/{sessionId}/live", wrapper.LiveStatus)
	})
	r.Group(func(r chi.Router) {
		r.Post("/session/{sessionId}/players", wrapper.JoinSession)
	})
	r.Group(func(r chi.Router) {
		r.Delete("/session/{sessionId}/poll", wrapper.StopPoll)
	})
	r.Group(func(r chi.Router) {
		r.Post("/session/{sessionId}/poll", wrapper.CreatePoll)
	})
	r.Group(func(r chi.Router) {
		r.Post("/session/{sessionId}/poll/vote", wrapper.Vote)
	})

	return r
}