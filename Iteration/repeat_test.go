package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("5")
	expected := "55555"

	if repeat != expected {
		t.Errorf("got repeat %q expected %q", repeat, expected)
	}
}
