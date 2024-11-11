package hello

import "testing"

var prefix string = "Hello, "

func Hello(name string) string {
	if name == "" {
		return prefix + "World"
	}
	return prefix + name
}

func TestHello(t *testing.T) {
	var name string = "yyj"
	got := Hello(name)
	want := prefix + name  
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}

	got = Hello("")
	want = "Hello" + ", World"
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}