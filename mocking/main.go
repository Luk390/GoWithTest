package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigureableSleeper struct {
	sleepDuration time.Duration
	sleep         func(time.Duration)
}

const countdownStart = 3
const finalWord = "Go"

func (s ConfigureableSleeper) Sleep() {
	s.sleep(s.sleepDuration)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)

}

func main() {
	sleeper := ConfigureableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
