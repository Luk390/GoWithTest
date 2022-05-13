package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	spySleeper := &SpySleeper{}
	buffer := &bytes.Buffer{}
	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := 
`3
2
1
Go`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if spySleeper.Calls != 4 {
		t.Errorf("expected spySleeper to be called 4 times, called %d times", spySleeper.Calls)
	}
}