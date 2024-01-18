package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	height float64
	base   float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.length + rectangle.breadth)
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return (t.base * t.height * 0.5)
}
