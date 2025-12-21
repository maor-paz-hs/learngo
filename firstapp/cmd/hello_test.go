//package main
//
//import (
//	"testing"
//)
//
//func assertCorrectMessage(t *testing.T, got, want string) {
//	t.Helper()
//	if got != want {
//		t.Errorf("got %q want %q", got, want)
//	}
//}
//
//func TestHello(t *testing.T) {
//	t.Run("saying hello to people", func(t *testing.T) {
//		got := Hello("Chris")
//		want := "Hello, Chris"
//		assertCorrectMessage := (t, got, want)
//	})
//	t.Run("empty string defaults to 'world'", func(t *testing.T) {
//		got := Hello("")
//		want := "Hello, World"
//		assertCorrectMessage := (t, got, want)
//	})
//}
//
//
//
//
//// func TestHelloFunction(t *testing.T) {
////	want := "Hello - World!"
////	if got := Hello(); got != want {
////		t.Errorf("Hello() = %q, want %q", got, want)
////	}
//// }
//
//// func TestMainFunction(t *testing.T) {
////	// This is a placeholder test for the main function.
////	// Since main() does not return a value, we cannot test it directly.
////	// However, we can ensure that it runs without panicking.
////	defer func() {
////		if r := recover(); r != nil {
////			t.Errorf("main() panicked: %v", r)
////		}
////	}()
////	main()
//// }
