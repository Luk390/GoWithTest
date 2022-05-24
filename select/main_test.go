package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("url from fast server returned", func(t *testing.T) {
		slowServer := createDelayedServer(10 * time.Millisecond)
		fastServer := createDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(fastUrl, slowUrl)

		if err != nil {
			t.Fatalf("got error when not expecting one %v", err)
		}
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})
	t.Run("timesout after 10 secs", func(t *testing.T) {
		serverA := createDelayedServer(25 * time.Millisecond)
		defer serverA.Close()

		u1 := serverA.URL
		u2 := serverA.URL

		_, err := configureableRacer(u1, u2, 20 * time.Millisecond)
		if err == nil {
			t.Errorf("Expected error but didn't get one")
		}
	})

}

func createDelayedServer(t time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
