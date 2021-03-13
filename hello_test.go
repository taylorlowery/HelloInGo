package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tater")
	want := "Hello, Tater"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
