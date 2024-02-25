package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "value"}

	t.Run("known word", func(t *testing.T) {
		got, err := dict.Search("test")
		want := "value"

		if err != nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
