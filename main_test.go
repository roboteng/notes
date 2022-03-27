package main_test

import (
	ts "notes/testHelpers"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "Given...",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			want := 2
			t.Run("When...", func(t *testing.T) {
				got := 1 + 1
				ts.AssertEquals(t, want, got)
			})
		})
	}
}
