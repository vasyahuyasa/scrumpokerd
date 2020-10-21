package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vasyahuyasa/scrumpokerd/api"
)

const port = 6777

func main() {
	scrumpocker := api.NewScrumpocker()

	h := api.Handler(scrumpocker)

	s := &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}

	log.Fatal(s.ListenAndServe())
}
