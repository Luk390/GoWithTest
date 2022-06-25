package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type spyStore struct {
	response string
	t        *testing.T
}

type spyResponseWriter struct {
	written bool
}

func (s *spyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *spyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *spyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *spyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				fmt.Println("got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello world"
		store := &spyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingContext, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingContext)

		response := &spyResponseWriter{}
		svr.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a response should not have been written")
		}
	})
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello world"
		store := &spyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s want %s", response.Body.String(), data)
		}

		//store.assertWasNotCancelled()
	})
}
