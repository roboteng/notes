package types

type Note struct {
	Title    string `json:"title"`
	Contents string `json:"desc"`
	Id       int    `json:"id"`
}

func (n *Note) Equals(other *Note) bool {
	return n.Title == other.Title && n.Contents == other.Contents
}

type NoteService interface {
	NotesViewer
	NoteCreator
	SingleNoteViewer
}

type NotesViewer interface {
	ViewNotes() []Note
}

type NoteCreator interface {
	CreateNote(note Note) (int, error)
}

type SingleNoteViewer interface {
	ViewSingleNote(id int) (Note, error)
}

type AnonNoteService struct {
	AnonNoteCreator
	AnonNotesViewer
	AnonSingleNoteViewer
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

type AnonSingleNoteViewer struct {
	View func(id int) (Note, error)
}

func (a *AnonSingleNoteViewer) ViewSingleNote(id int) (Note, error) {
	return a.View(id)
}

type CreateNoteResponse struct {
	Id int `json:"id"`
}
