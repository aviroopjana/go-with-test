package main

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	expected := 16

	if got != expected {
		t.Errorf("got %d expected %d, given %v", got, expected, numbers)
	}

}
