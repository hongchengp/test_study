package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaa"
	if repeated != expected {
		t.Fatalf("expected %s, got %s", repeated, expected)
	}
}
