package main

import "testing"

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{20, 20}
		got := rectangle.Area()
		expected := 400.0

		if got != expected {
			t.Errorf("got %g expected %g", got, expected)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		expected := 314.1592653589793

		if got != expected {
			t.Errorf("got %g expected %g", got, expected)
		}
	})
}
