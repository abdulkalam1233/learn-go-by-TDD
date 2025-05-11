package maps

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "Test something"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "Test something"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("Expected an error")
		}
		assertErrors(t, got, ErrNotFound)
	})
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrors(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
