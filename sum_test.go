package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7, 9}

		got := Sum(numbers)
		expected := 25

		assertCorrectInt(t, got, expected)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{2, 4, 6}

		got := Sum(numbers)
		expected := 12

		assertCorrectInt(t, got, expected)
	})
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 3, 5}
	numbers2 := []int{2, 4, 6}

	got := SumAll(numbers1, numbers2)
	expected := []int{9, 12}

	assertCorrectVal(t, got, expected)
}

func TestSumAllTails(t *testing.T) {
	numbers1 := []int{1, 3}
	numbers2 := []int{2, 4, 6}
	numbers3 := []int{3, 5, 7}

	got := SumAllTails(numbers1, numbers2, numbers3)
	expected := []int{3, 10, 12}

	assertCorrectVal(t, got, expected)
}

func assertCorrectInt(t *testing.T, got, expected int) {
	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}

func assertCorrectVal(t *testing.T, got, expected []int) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected '%v' but got '%v'", expected, got)
	}
}
