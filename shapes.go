package main

import "math"

type Rectangle struct {
	x float64
	y float64
}

type Circle struct {
	r float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.x + r.y)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r Rectangle) Area() float64 {
	return r.x * r.y
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}
