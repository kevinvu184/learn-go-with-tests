package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Provide both name and locale", func(t *testing.T) {
		got := hello("Kevin", "en-GB")
		want := "Hello Kevin!"

		assertToEqual(t, got, want)
	})

	t.Run("Provide locale, but not name", func(t *testing.T) {
		got := hello("", "en-GB")
		want := "Hello World!"

		assertToEqual(t, got, want)
	})

	t.Run("Provide both name and locale in es-ES", func(t *testing.T) {
		got := hello("Kevin", "es-ES")
		want := "Hola Kevin!"

		assertToEqual(t, got, want)
	})

	t.Run("Provide none of name and locale", func(t *testing.T) {
		got := hello("", "")
		want := "Hello World!"

		assertToEqual(t, got, want)
	})

	t.Run("Provide locale, but not name in es-ES`", func(t *testing.T) {
		got := hello("", "es-ES")
		want := "Hola World!"

		assertToEqual(t, got, want)
	})

	t.Run("Provide both name and locale in fr-FR", func(t *testing.T) {
		got := hello("Kevin", "fr-FR")
		want := "Bonjour Kevin!"

		assertToEqual(t, got, want)
	})
}

func assertToEqual(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
