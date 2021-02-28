package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
