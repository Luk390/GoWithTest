package main

import "testing"

func TestCounter (t *testing.T) {
	t.Run("increment counter 3 times returns 3", func(t *testing.T){
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		got := counter.Value()
		want := 3

		if got != want {
			t.ErrorF("got %v want %v", got, want)
		}
	})
}