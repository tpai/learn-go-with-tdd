package main

import (
	"log"
	"net/http"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	playTimes map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.playTimes[name]
}

func (i *InMemoryPlayerStore) RecordSave(name string) {
	i.playTimes[name]++
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
