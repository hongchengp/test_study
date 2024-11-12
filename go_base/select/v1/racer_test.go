package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer fastServer.Close()

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 *time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer slowServer.Close()
	fastUrl := fastServer.URL
	slowUrl := slowServer.URL
	got := Racer(slowUrl, fastUrl)
	want := fastUrl 
	if got != want {
		t.Fatalf("want %s, but got %s", want, got)
	}
}

