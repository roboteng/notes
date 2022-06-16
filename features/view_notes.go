package features

import "notes/types"

type ViewNotes struct {
	Service types.NotesViewer
}

func (v *ViewNotes) View() ([]types.Note, error) {
	return v.Service.View(), nil
}
