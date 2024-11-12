package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	CountDown(buffer)
	got := buffer.String()
	want := "3"
	if got != want {
		t.Fatalf("want %s, but got %s", want, got)
	}
}