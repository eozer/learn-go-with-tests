package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	// dictionary := map[string]string{"test": "foobar"}
	d := Dictionary{"test": "foobar"}
	t.Run("known word", func(t *testing.T) {
		got, _ := d.Search("test")
		want := "foobar"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := d.Search("nope")
		want := "not found"
		if err == nil {
			t.Fatal("Expected to get an error")
		}
		assertStrings(t, err.Error(), want)
	})

	t.Run("add, update, and delete word", func(t *testing.T) {
		_, err := d.Search("nope")
		want := "not found"
		if err == nil {
			t.Fatal("Expected to get an error")
		}
		assertStrings(t, err.Error(), want)
		//
		d.Add("nope", "nope definition")
		got, _ := d.Search("nope")
		fmt.Println(got)
		assertStrings(t, got, "nope definition")
		// Search again
		_, err = d.Search("nope")
		if err != nil {
			t.Fatal("Expected to NOT get an error")
		}
		d.Update("nope", "nope2")
		got, _ = d.Search("nope")
		assertStrings(t, got, "nope2")
		//
		d.Delete("nope")
		_, err = d.Search("nope")
		if err == nil {
			t.Fatal("Expected to get an error")
		}
	})
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}

}
