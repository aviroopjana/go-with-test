package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(20.0, 20.0)
	expected := 80.0

	if got != expected {
		t.Errorf("got %f expected %f", got, expected)
	}
}

func TestArea(t *testing.T) {
	got := Area(20.0, 20.0)
	expected := 400.0

	if got != expected {
		t.Errorf("got %f expected %f", got, expected)
	}
}
