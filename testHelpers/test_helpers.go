package testhelpers

import "testing"

func AssertEquals[T comparable](t testing.TB, want, got T, flavor string) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %v %v, but got %v", flavor, want, got)
	}
}

func AssertEqualSlice[T comparable](t testing.TB, want, got []T) {
	t.Helper()
	AssertEquals(t, len(want), len(got), "length")
	for i, v := range want {
		AssertEquals(t, v, got[i], "")
	}
}
