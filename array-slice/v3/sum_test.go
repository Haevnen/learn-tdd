package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertCorrectMessage(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("2 collection", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		// assertCorrectSlice(t, got, want)
		// seeing if any two variables are the same
		// not type safe
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("1 collection only", func(t *testing.T) {
		got := SumAll([]int{1, 2})
		want := []int{3}

		assertCorrectSlice(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertCorrectSlice(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("got %v want %v", got, want)
	}

	for i, val := range got {
		if val != want[i] {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
