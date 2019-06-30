package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSavingScoresAndRetrieveThem(t *testing.T) {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	player := "Tony"

	server.ServeHTTP(httptest.NewRecorder(), postScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), postScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), postScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, getScoreRequest(player))

	assertResponseStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
