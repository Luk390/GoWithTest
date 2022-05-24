package main

import (
	"fmt"
	"net/http"
	"time"
)

const timeout = 10 * time.Second

func Racer(u1, u2 string) (winner string, err error) {
	return configureableRacer(u1, u2, timeout)
}

func configureableRacer(u1, u2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(u1):
		return u1, nil
	case <-ping(u2):
		return u2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("no response from %s and %s", u1, u2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}