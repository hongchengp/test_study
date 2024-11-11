package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Fatalf("expected %s, got %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 5)
	}
}