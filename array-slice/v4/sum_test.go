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
	assertCorrectSlice := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

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

func TestSumAllTails(t *testing.T) {
	assertCorrectSlice := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make the sum of tails of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assertCorrectSlice(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
