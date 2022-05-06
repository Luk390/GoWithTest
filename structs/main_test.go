package main

import "testing"

func TestPerimeter(t *testing.T) {
	width := 10.2
	height := 2.0
	got := Perimeter(width, height)
	want := 20.4 + 4.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

