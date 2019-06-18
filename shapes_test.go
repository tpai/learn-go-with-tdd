package main

import "math"
import "testing"

type Shape interface {
	Perimeter() float64
	Area() float64
}

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.0, 5.0}
		expected := 20.0
		checkPerimeter(t, rectangle, expected)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{5.0}
		expected := 2 * math.Pi * 5.0
		checkPerimeter(t, circle, expected)
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{3.0, 5.0}
		expected := 15.0
		checkArea(t, rectangle, expected)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{5.0}
		expected := math.Pi * 25.0
		checkArea(t, circle, expected)
	})
}

func checkPerimeter(t *testing.T, shape Shape, expected float64) {
	t.Helper()
	got := shape.Perimeter()
	if got != expected {
		t.Errorf("expected '%.2f' but got '%.2f'", expected, got)
	}
}

func checkArea(t *testing.T, shape Shape, expected float64) {
	t.Helper()
	got := shape.Area()
	if got != expected {
		t.Errorf("expected '%.2f' but got '%.2f'", expected, got)
	}
}
