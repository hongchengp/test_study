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
			[]struct {
				Name string
				Age int
			} {
				{
					"yyj",
					21,
				},
				{
					"hcp", 
					22,
				},
			},
			[]string{"yyj", "hcp"},
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

	t.Run("with maps", func(t *testing.T) {
    aMap := map[string]string{
        "Foo": "Bar",
        "Baz": "Boz",
    }

    var got []string
    walk(aMap, func(input string) {
        got = append(got, input)
    })

    assertContains(t, got, "Bar")
    assertContains(t, got, "Boz")
})
}


func assertContains(t *testing.T, haystack []string, needle string)  {
    contains := false
    for _, x := range haystack {
        if x == needle {
            contains = true
        }
    }
    if !contains {
        t.Errorf("expected %+v to contain '%s' but it didnt", haystack, needle)
    }
}