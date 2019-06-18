package main

import "math"
import "testing"

type Shape interface {
	Perimeter() float64
	Area() float64
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{5.0, 5.0}, 20.0},
		{Circle{5.0}, 2 * math.Pi * 5.0},
		{Triangle{3.0, 4.0}, 9.0},
	}
	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.expected {
			t.Errorf("%#v expected '%.2f' but got '%.2f'", tt.shape, tt.expected, got)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{3.0, 5.0}, 15.0},
		{Circle{5.0}, math.Pi * 25.0},
		{Triangle{3.0, 4.0}, 6.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.expected {
			t.Errorf("%#v expected '%.2f' but got '%.2f'", tt.shape, tt.expected, got)
		}
	}
}
