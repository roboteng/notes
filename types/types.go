package types

type Note struct {
	Title    string `json:"title"`
	Contents string `json:"desc"`
	Id       int    `json:"id"`
}

type NoteService interface {
	NotesViewer
	NoteCreator
}

type NotesViewer interface {
	ViewNotes() []Note
}

type NoteCreator interface {
	CreateNote(note Note) (int, error)
}

type AnonNoteService struct {
	AnonNoteCreator
	AnonNotesViewer
}

type AnonNotesViewer struct {
	View func() []Note
}

func (a *AnonNotesViewer) ViewNotes() []Note {
	return a.View()
}

type AnonNoteCreator struct {
	Create func(note Note) (int, error)
}

func (a *AnonNoteCreator) CreateNote(note Note) (int, error) {
	return a.Create(note)
}
