package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Channel struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
}

type Message struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	Sender string `json:"sender"`
}

var channels = []Channel{
	{
		ID:    1,
		Label: "general",
	},
	{
		ID:    2,
		Label: "random",
	},
}

var messages = []Message{
	{
		ID:     1,
		Body:   "Hello world!",
		Sender: "Alice",
	},
	{
		ID:     2,
		Body:   "Hello to you, too!",
		Sender: "Bob",
	},
}

func handleListChannels(w http.ResponseWriter, r *http.Request) {
	bytes, _ := json.Marshal(channels)
	w.Write(bytes)
}

func handleListMessages(w http.ResponseWriter, r *http.Request) {
	bytes, _ := json.Marshal(messages)
	w.Write(bytes)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/channels", handleListChannels)
	r.Get("/api/messages/{channelId}", handleListMessages)

	log.Print("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
