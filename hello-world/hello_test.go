package hello_world

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Say 'Hello, world' when name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Greet in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
