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
	View() []Note
}

type NoteCreator interface {
	Save(note Note) (int, error)
}

type SingleNoteViewer interface {
	ViewSingle(id int) (Note, error)
}

type AnonNoteService struct {
	AnonNoteCreator
	AnonNotesViewer
	AnonSingleNoteViewer
}

type AnonNotesViewer struct {
	ViewNotes func() []Note
}

func (a *AnonNotesViewer) View() []Note {
	return a.ViewNotes()
}

type AnonNoteCreator struct {
	SaveNote func(note Note) (int, error)
}

func (a *AnonNoteCreator) Save(note Note) (int, error) {
	return a.SaveNote(note)
}

type AnonSingleNoteViewer struct {
	View func(id int) (Note, error)
}

func (a *AnonSingleNoteViewer) ViewSingle(id int) (Note, error) {
	return a.View(id)
}

type CreateNoteResponse struct {
	Id int `json:"id"`
}
