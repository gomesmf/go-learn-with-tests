package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	sayHelloPeople := func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, " + "Chriss"
		assertCorrectMessage(t, got, want)
	}
	sayHelloWorld := func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	}
	t.Run("say hello to people", sayHelloPeople)
	t.Run("say hello to world when empty string is provided", sayHelloWorld)
}
