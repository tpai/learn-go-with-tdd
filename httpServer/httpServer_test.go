package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores    map[string]int
	saveCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordSave(name string) {
	s.saveCalls = append(s.saveCalls, name)
}

func TestGetScoreByPlayer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Tony":     1988,
			"Tigerfat": 1987,
		},
		[]string{},
	}
	server := &PlayerServer{store: &store}

	t.Run("returns Tony's score", func(t *testing.T) {
		request := getScoreRequest("Tony")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "1988")
	})
	t.Run("returns Tigerfat's score", func(t *testing.T) {
		request := getScoreRequest("Tigerfat")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "1987")
	})
	t.Run("returns 404 on missing player", func(t *testing.T) {
		request := getScoreRequest("Johndoe")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		expected := http.StatusNotFound

		assertResponseStatus(t, got, expected)
	})
}

func getScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return req
}

func assertResponseStatus(t *testing.T, got, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("status code is wrong, expected '%d' but got '%d'", expected, got)
	}
}

func assertResponseBody(t *testing.T, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("response body is wrong, expected '%s' but got '%s'", expected, got)
	}
}

func TestPostScoreToStore(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Tony"

		request := postScoreRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusAccepted)

		if len(store.saveCalls) != 1 {
			t.Errorf("save score calls is wrong, expected %d but got %d", 1, len(store.saveCalls))
		}

		if store.saveCalls[0] != player {
			t.Errorf("did not store correct player, expected %s but got %s", player, store.saveCalls[0])
		}
	})
}

func postScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}
