package services

import (
	"errors"
	"notes/types"
)

type InMemoryNoteService struct {
	notes []types.Note
}

func NewInMemoryNoteService() *InMemoryNoteService {
	return &InMemoryNoteService{
		notes: make([]types.Note, 0),
	}
}

func (i *InMemoryNoteService) CreateNote(note types.Note) (int, error) {
	id := len(i.notes) + 1
	note.Id = id
	i.notes = append(i.notes, note)
	return id, nil
}

func (i *InMemoryNoteService) ViewNotes() []types.Note {
	return i.notes
}

func (i *InMemoryNoteService) ViewSingleNote(id int) (types.Note, error) {
	if id > len(i.notes) {
		return types.Note{}, errors.New("No note found with that ID")
	}
	note := i.notes[id-1]
	return note, nil
}
