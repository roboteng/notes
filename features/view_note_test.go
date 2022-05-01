package features_test

import (
	"notes/features"
	testhelpers "notes/testHelpers"
	"notes/types"
	"testing"
)

func TestViewNotes(t *testing.T) {
	t.Run("Takes a service", func(t *testing.T) {
		service := &types.AnonNotesViewer{
			View: func() []types.Note {
				return []types.Note{}
			},
		}
		viewNotes := features.ViewNotes{
			Service: service,
		}
		notes, err := viewNotes.View()
		if err != nil {
			t.Errorf("Got unexpected error: %v", err)
			return
		}
		testhelpers.AssertEquals(t, notes, []types.Note{}, "Notes list")
	})
}
