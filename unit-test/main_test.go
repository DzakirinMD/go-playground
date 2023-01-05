package main

import (
	"fmt"
	"testing"
)

// run go test -v
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

func TestSubtractTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{10, 1, 9},
		{1, 0, 1},
		{3, 2, 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// t.Run enables running “subtests”, one for each table entry. These are shown separately when executing go test -v.
		t.Run(testname, func(t *testing.T) {
			ans := Subtract(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Basic test
func TestIntMinBasic(t *testing.T) {
	got := IntMin(2, -2)
	want := -2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// Writing tests can be repetitive, so it’s idiomatic to use a table-driven style,
// where test inputs and expected outputs are listed in a table and a single loop walks over them and performs the test logic.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// t.Run enables running “subtests”, one for each table entry. These are shown separately when executing go test -v.
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmarking test
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
