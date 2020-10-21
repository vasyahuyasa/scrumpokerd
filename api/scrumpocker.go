package api

import (
	"fmt"
	"log"
	"net/http"
)

type Scrumpocker struct {
	broker *broker
}

type broker struct {
	// Events are pushed to this channel by the main events-gathering routine
	notifier chan []byte

	// New client connections
	newClients chan chan []byte

	// Closed client connections
	closingClients chan chan []byte

	// Client connections registry
	clients map[chan []byte]bool
}

func NewScrumpocker() *Scrumpocker {
	b := &broker{
		notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        map[chan []byte]bool{},
	}

	b.listen()

	return &Scrumpocker{
		broker: b,
	}
}

func (scrumpocker *Scrumpocker) CreateSession(w http.ResponseWriter, r *http.Request) {

}

func (scrumpocker *Scrumpocker) GetSession(w http.ResponseWriter, r *http.Request, sessionId string) {

}

func (scrumpocker *Scrumpocker) LiveStatus(w http.ResponseWriter, r *http.Request, sessionId string) {
	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan []byte)

	// Signal the broker that we have a new connection
	scrumpocker.broker.newClients <- messageChan

	// Remove this client from the map of connected clients
	// when this handler exits.
	defer func() {
		scrumpocker.broker.closingClients <- messageChan
	}()

	select {
	case <-r.Context().Done():
		break

	case msg := <-messageChan:
		_, err := fmt.Fprintf(w, "data: %s\n\n", msg)
		if err != nil {
			log.Printf("can not send message to client: %v", err)
		}
		flusher.Flush()
	}
}

func (scrumpocker *Scrumpocker) JoinSession(w http.ResponseWriter, r *http.Request, sessionId string) {

}

func (scrumpocker *Scrumpocker) StopPoll(w http.ResponseWriter, r *http.Request, sessionId string) {

}

func (scrumpocker *Scrumpocker) CreatePoll(w http.ResponseWriter, r *http.Request, sessionId string) {

}

func (scrumpocker *Scrumpocker) Vote(w http.ResponseWriter, r *http.Request, sessionId string) {

}

func (broker *broker) listen() {
	for {
		select {
		case s := <-broker.newClients:
			// A new client has connected.
			// Register their message channel
			broker.clients[s] = true
			log.Printf("Client added. %d registered clients", len(broker.clients))

		case s := <-broker.closingClients:
			// A client has dettached and we want to
			// stop sending them messages.
			delete(broker.clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.clients))

		case event := <-broker.notifier:
			for clientMessageChan := range broker.clients {
				clientMessageChan <- event
			}
		}
	}

}
