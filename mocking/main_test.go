package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type spyTime struct {
	sleepDuration time.Duration
}

func (s *spyTime) Sleep(duration time.Duration) {
	s.sleepDuration = duration
}

func TestConfigureableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spy := &spyTime{}
	sleeper := ConfigureableSleeper{
		sleepDuration: sleepTime,
		sleep:         spy.Sleep,
	}
	sleeper.Sleep()

	if sleeper.sleepDuration != spy.sleepDuration {
		t.Errorf("unexpected sleepDuration got %v want %v", spy.sleepDuration, sleepTime)
	}
}

func TestCountdown(t *testing.T) {
	t.Run("prints 321Go", func(t *testing.T) {

		SpyCountdownOperations := &SpyCountdownOperations{}
		buffer := &bytes.Buffer{}
		Countdown(buffer, SpyCountdownOperations)

		got := buffer.String()
		want :=
			`3
2
1
Go`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleeps before every print", func(t *testing.T) {
		spyCountdownOperations := &SpyCountdownOperations{}
		Countdown(spyCountdownOperations, spyCountdownOperations)
		got := spyCountdownOperations.Calls
		want := []string{"sleep", "write", "sleep", "write", "sleep", "write", "sleep", "write"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("call order: got %v want %v", got, want)
		}
	})
}
