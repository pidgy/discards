package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/pidgy/discards/battle"
	"github.com/pidgy/discards/options"
)

func main() {
	api := flag.String("api", ".api", "path to api key file")
	flag.Parse()

	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.Stamp}).With().Timestamp().Logger()

	key, err := os.ReadFile(*api)
	if err != nil {
		panic(err)
	}
	options.APIKey = string(key)

	addr := "localhost:8080"

	log.Info().Str("addr", addr).Msg("discards server started")

	http.HandleFunc("/card", card)
	http.HandleFunc("/sets", sets)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func card(w http.ResponseWriter, r *http.Request) {
	log.Info().Str("from", r.RemoteAddr).Str("type", "request").Msg(r.RequestURI)

	final := log.Info().Str("from", r.RemoteAddr).Str("type", "response")
	defer final.Msg(r.RequestURI)

	if r.Method != "GET" {
		final.Int("status", http.StatusNotFound)
		http.Error(w, "method is not supported.", http.StatusNotFound)
		return
	}

	id := r.URL.Query().Get("id")

	c := &battle.Card{}
	err := c.Get(id)
	if err != nil {
		final.Err(err).Int("status", http.StatusNotFound)
		http.Error(w, "failed to find a card with id: "+id, http.StatusNotFound)
		return
	}

	final.Str("card", c.Name)

	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		final.Err(err).Int("status", http.StatusInternalServerError)
		http.Error(w, "failed to send encoded card: "+id, http.StatusInternalServerError)
		return
	}
}

func sets(w http.ResponseWriter, r *http.Request) {
	log.Info().Str("from", r.RemoteAddr).Msg(r.RequestURI)

	final := log.Info().Str("from", r.RemoteAddr)
	defer final.Msg(r.RequestURI)

	if r.Method != "GET" {
		final.Int("status", http.StatusNotFound)
		http.Error(w, "method is not supported.", http.StatusNotFound)
		return
	}

	s := &battle.Sets{}
	err := s.Get()
	if err != nil {
		final.Err(err).Int("status", http.StatusNotFound)
		http.Error(w, "failed to find any sets", http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		final.Err(err).Int("status", http.StatusInternalServerError)
		http.Error(w, "failed to send encoded sets.", http.StatusInternalServerError)
		return
	}
}
