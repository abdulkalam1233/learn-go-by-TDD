package maps

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dictionary := map[string]string{"test": "Test something"}
	got := Search(dictionary, "test")
	want := "Test something"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
