package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Tony", "English")
		want := "Hello Tony"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when nothing input", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("Tony", "Chinese")
		want := "你好 Tony"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Tony", "Japanese")
		want := "こんにちは Tony"

		assertCorrectMessage(t, got, want)
	})
}
