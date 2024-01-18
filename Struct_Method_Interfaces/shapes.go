package main

import "math"

type Rectangle struct {
	length  float64
	breadth float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
