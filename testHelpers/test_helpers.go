package testhelpers

import "testing"

func AssertEquals[T comparable](t testing.TB, want, got T) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}
