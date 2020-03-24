package iterator

import "testing"

func TestRepeat(t *testing.T) {
	want := "AAAA"
	got := Repeat("A", 4)
	if got != want {
		t.Errorf("Wanted %s, got %s", want, got)
	}
}
