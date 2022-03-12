package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/pidgy/discards/battle"
	"github.com/pidgy/discards/options"
)

func main() {
	key, err := os.ReadFile(".api")
	if err != nil {
		panic(err)
	}
	options.APIKey = string(key)

	fmt.Printf("Starting discards server at port 8080\n")

	http.HandleFunc("/card", card)
	http.HandleFunc("/sets", sets)

	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}

func card(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/card" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	id := r.URL.Query().Get("id")

	c := &battle.Card{}
	err := c.Get(id)
	if err != nil {
		http.Error(w, "Failed to find card with id \""+id+"\".", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		http.Error(w, "Failed to send encoded data.", http.StatusInternalServerError)
		return
	}
}

func sets(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sets" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	s := &battle.Sets{}
	err := s.Get()
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		http.Error(w, "Failed to send encoded data.", http.StatusInternalServerError)
		return
	}
}
