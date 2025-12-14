package main

import (
	"testing"
)

func TestHelloFunction(t *testing.T) {
	want := "Hello - World!"
	if got := Hello(); 
	got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}