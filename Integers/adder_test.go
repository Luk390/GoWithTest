package main

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	got := Adder(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleAdder() {
	sum := Adder(1, 2)
	fmt.Println(sum)
	// Output:
	// 3
}
