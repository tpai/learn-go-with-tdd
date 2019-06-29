package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leave it at 3", func(t *testing.T) {
		counter := &Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		expected := 1000
		counter := &Counter{}

		var wg sync.WaitGroup
		wg.Add(expected)

		for i := 0; i < expected; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, expected)
	})
}

func assertCounter(t *testing.T, got *Counter, expected int) {
	t.Helper()
	if got.Value() != expected {
		t.Errorf("expected %d but got %d", expected, got.Value())
	}
}
