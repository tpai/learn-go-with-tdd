package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("should return fast url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		expected := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != expected {
			t.Errorf("expected '%s' but got '%s'", expected, got)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10 seconds", func(t *testing.T) {
		slowServerA := makeDelayedServer(11 * time.Second)
		slowServerB := makeDelayedServer(12 * time.Second)

		_, err := ConfigurableRacer(slowServerA.URL, slowServerB.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but got nothing")
		}
	})
}

func makeDelayedServer(delayedTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayedTime)
		w.WriteHeader(http.StatusOK)
	}))
}
