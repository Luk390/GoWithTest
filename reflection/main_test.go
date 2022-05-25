package main

import "testing"

func TestWalk(t *testing.T) {
	expected := "Luke"
	var got []string
	x := struct {
		name string
	}{expected}
	walk(x, func(input string) {
		got = append(got, input)
	})
	if len(got) != 1 {
		t.Errorf("wrong number of function calls got %v want %v", len(got), 1)
	}
	if got[0] != expected {
		t.Errorf("got %s want %s", got[0], expected)
	}
}
