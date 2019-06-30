package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordSave(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.saveScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	result := p.store.GetPlayerScore(player)

	if result == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, result)
}

func (p *PlayerServer) saveScore(w http.ResponseWriter, player string) {
	p.store.RecordSave(player)

	w.WriteHeader(http.StatusAccepted)
}
