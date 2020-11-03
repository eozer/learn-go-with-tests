package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedHTTPServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"
	t.Run("old way", func(t *testing.T) {
		got := Racer(slowURL, fastURL)
		want := fastURL
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("mock way", func(t *testing.T) {
		slowServer := makeDelayedHTTPServer(1 * time.Second)
		fastServer := makeDelayedHTTPServer(100 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		got := Racer(slowURL, fastURL)
		want := fastURL
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
