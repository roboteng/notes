package types

type Note struct {
	Title    string `json:"title"`
	Contents string `json:"desc"`
}

type NotesViewer interface {
	ViewNotes() []Note
}

type AnonNotesViewer struct {
	View func() []Note
}

func (a *AnonNotesViewer) ViewNotes() []Note {
	return a.View()
}
