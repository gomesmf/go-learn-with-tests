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
