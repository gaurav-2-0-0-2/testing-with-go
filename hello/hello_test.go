package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Eloide", "Spanish")
		want := "Hola, Eloide"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Eloide", "French")
		want := "Bonjour, Eloide"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Sanskrit", func(t *testing.T) {
		got := Hello("Eloide", "Sanskrit")
		want := "नमस्कारः, Eloide"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Eloide", "German")
		want := "begrüßen, Eloide"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// to tell the test suit that this method is a helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
