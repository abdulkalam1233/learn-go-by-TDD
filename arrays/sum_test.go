package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		given := []int{1, 2, 3, 4, 5}
		got := Sum(given)
		want := 15
		assertCorrectMessage(t, got, want, given)
	})

	t.Run("Collection of any size", func(t *testing.T) {
		given := []int{1, 2, 3}
		got := Sum(given)
		want := 6
		assertCorrectMessage(t, got, want, given)
	})
}

func assertCorrectMessage(t *testing.T, got int, want int, given []int) {
	if got != want {
		t.Errorf("got %d, want %d given %v", got, want, given)
	}
}
