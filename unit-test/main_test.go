package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(10, 10)
	want := 20

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(20, 10)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
