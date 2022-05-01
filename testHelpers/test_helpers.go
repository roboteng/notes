package testhelpers

import (
	"reflect"
	"testing"
)

func AssertEquals[T any](t testing.TB, want, got T, flavor string) bool {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v %v, but got %v", flavor, want, got)
		return false
	}
	return true
}

func AssertEqualSlice[T comparable](t testing.TB, want, got []T) bool {
	t.Helper()
	if !AssertEquals(t, len(want), len(got), "length") {
		return false
	}
	for i, v := range want {
		AssertEquals(t, v, got[i], "")
		return false
	}
	return true
}
