package main

import "testing"

func TestPrintsHelloWorld(t *testing.T) {

	assertEquals := func(t *testing.T, expected, actual string) {
		// Skip line information when a test fails.
		t.Helper()
		if actual != expected {
			t.Errorf("expected %q got %q", expected, actual)
		}
	}

	t.Run("it prints hello world when empty string is given", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello World!"
		assertEquals(t, expected, actual)
	})

	t.Run("it prints hello {name} when {name} is given", func(t *testing.T) {
		actual := Hello("Suleyman", "")
		expected := "Hello Suleyman!"
		assertEquals(t, expected, actual)
	})

	t.Run("it prints in Spanish", func(t *testing.T) {
		actual := Hello("Suleyman", "es")
		expected := "Hola Suleyman!"
		assertEquals(t, expected, actual)
	})

	t.Run("it prints in French", func(t *testing.T) {
		actual := Hello("Suleyman", "fr")
		expected := "Bonjour Suleyman!"
		assertEquals(t, expected, actual)
	})
}
