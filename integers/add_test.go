package integers

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Add two positive numbers", func(t *testing.T) {
		expected := Add(1, 2)
		got := 3
		assertCorrectMessage(t, expected, got)
	})

}

func assertCorrectMessage(t *testing.T, expected int, got int) {
	t.Helper() // This will help to identify the test that failed
	if expected != got {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}
