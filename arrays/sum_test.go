package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		given := [5]int{1, 2, 3, 4, 5}
		got := Sum(given)
		want := 15
		assertCorrectMessage(t, got, want, given)
	})
}

func assertCorrectMessage(t *testing.T, got int, want int, given [5]int) {
	if got != want {
		t.Errorf("got %d, want %d given %v", got, want, given)
	}
}
