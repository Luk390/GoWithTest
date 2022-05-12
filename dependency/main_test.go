package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Luke")
	got := buffer.String()
	want := "Hello, Luke"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}