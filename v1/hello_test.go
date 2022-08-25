package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Arpit")
	want := "Hello, Arpit"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
