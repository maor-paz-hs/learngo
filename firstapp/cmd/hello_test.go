package main

import (
	"testing"
)

func TestHelloFunction(t *testing.T) {
	want := "Hello - World!"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestMainFunction(t *testing.T) {
	// This is a placeholder test for the main function.
	// Since main() does not return a value, we cannot test it directly.
	// However, we can ensure that it runs without panicking.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main() panicked: %v", r)
		}
	}()
	main()
}
