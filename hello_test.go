package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Douglas")
	want := "Hello, Douglas"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
