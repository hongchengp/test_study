package main

import "testing"

func TestWalk(t *testing.T) {
	want := "yyj"
	var got []string
	x := struct {
		Name string
	} {
		Name: "yyj",
	}
	walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != want {
		t.Fatalf("want %s, but got %s", want, got)
	}
}