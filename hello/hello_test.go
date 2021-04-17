package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mohamed")
	want := "Hello, Mohamed"

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
