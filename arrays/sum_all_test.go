package arrays

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("Addition of all numbers in an array", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 3})
		want := []int{3, 6}
		checkSums(t, got, want)
	})
}

func TestSumAllAtTail(t *testing.T) {
	//checkSums := func(t *testing.T, got, want []int) {
	//	t.Helper()
	//	if !reflect.DeepEqual(got, want) {
	//		t.Errorf("got %v, want %v", got, want)
	//	}
	//}

	t.Run("Addition of numbers at tail", func(t *testing.T) {
		got := SumAllAtTail([]int{1, 2}, []int{3, 3})
		want := []int{2, 3}
		checkSums(t, got, want)
	})

	t.Run("Addition of empty array", func(t *testing.T) {
		got := SumAllAtTail([]int{}, []int{3, 3})
		want := []int{0, 3}
		checkSums(t, got, want)
	})

}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
