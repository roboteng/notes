package features_test

import (
	"notes/features"
	ts "notes/testHelpers"
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
		ts.AssertEquals(t, notes, []types.Note{})
	})
}
