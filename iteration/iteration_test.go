package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat character 5 times ", func(t *testing.T) {
		got := Repeat("a")
		expected := "aaaaa"
		assertCorrectMessage(t, expected, got)
	})
}

func assertCorrectMessage(t *testing.T, expected string, got string) {
	t.Helper() // This will help to identify the test that failed
	if expected != got {
		t.Errorf("expected '%s' but got '%s'", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
