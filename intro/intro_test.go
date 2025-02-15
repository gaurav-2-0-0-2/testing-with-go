package intro

import "testing"

func TestIntro(t *testing.T) {
	t.Run("if no name is passed", func(t *testing.T) {
		got := GiveIntro("", 4)
		want := "Name is required."
		assertMessage(t, got, want)
	})

	t.Run("if no class is passed", func(t *testing.T) {
		got := GiveIntro("Jhon", 0)
		want := "Class can't be zero"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}


