package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer)
	got := buffer.String()
	want := `3
2
1
Go!`
	if want != got {
		t.Fatalf("want %s, but got %s", want, got)
	}

}