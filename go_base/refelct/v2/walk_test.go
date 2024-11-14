package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	tests := []struct {
		Name string
		Calls interface{}
		Expected []string
	} {
		{
			"struct with String",
			struct {
				Name string
				Age int
			} {
				"yyj",
				21,
			},
			[]string{"yyj"},
		},
	}

	for _, test := range tests {
		var got []string
		walk(test.Calls, func(input string) {
			got = append(got, input)
		})

		want := test.Expected
		if ! reflect.DeepEqual(got, want) {
			t.Fatalf("want %v, but got %v", want, got)
		}
	}
}