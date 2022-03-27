package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertSum := func(t testing.TB, got int, expected int, numbers []int) {
		t.Helper()
		if got != expected {
			t.Errorf("got %d expected %d, given %v", got, expected, numbers)
		}
	}
	sumAnySize := func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		assertSum(t, got, expected, numbers)
	}
	t.Run("collection of any size", sumAnySize)
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(t, got, expected)
	})
}
