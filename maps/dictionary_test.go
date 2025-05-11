package maps

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "Test something"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "Test something"
		if err != nil {
			t.Fatal("should find added word:", err)
		}
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

func TestAdd(t *testing.T) {

	t.Run("Adding a word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "Test something"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Adding a word that already exists", func(t *testing.T) {
		word := "test"
		definition := "Test something"
		dictionary := Dictionary{word: definition}
		newDefinition := "New definition"
		err := dictionary.Add(word, newDefinition)
		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word string, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
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
