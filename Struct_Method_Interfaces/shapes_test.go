package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{20.0, 20.0}
	got := Perimeter(rectangle)
	expected := 80.0

	if got != expected {
		t.Errorf("got %f expected %f", got, expected)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{20.0, 20.0}
	got := Area(rectangle)
	expected := 400.0

	if got != expected {
		t.Errorf("got %f expected %f", got, expected)
	}
}
