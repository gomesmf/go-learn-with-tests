package arrays

import "testing"

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
