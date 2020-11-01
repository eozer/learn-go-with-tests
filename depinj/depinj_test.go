package main

import (
	"bytes"
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("initial test", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Foobar")
		got := buffer.String()
		want := "Hello Foobar"
		assertStrings(t, got, want)
	})

}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}

}
