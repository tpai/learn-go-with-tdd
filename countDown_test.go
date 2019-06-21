package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountDownOperationSpy struct {
	Calls []string
}

const sleep = "sleep"
const write = "write"

func (s *CountDownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountDownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountDown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		CountDown(buffer, spySleeper)

		got := buffer.String()
		expected := `3
2
1
GO!`

		if got != expected {
			t.Errorf("expected '%s' but got '%s'", expected, got)
		}

		if spySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got '%d'", spySleeper.Calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &CountDownOperationSpy{}

		CountDown(spySleeperPrinter, spySleeperPrinter)

		expected := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expected, spySleeperPrinter.Calls) {
			t.Errorf("expected calls '%v' got '%v'", expected, spySleeperPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
