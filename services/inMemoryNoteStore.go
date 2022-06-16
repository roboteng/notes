package services

import (
	"errors"
	"notes/types"
)

type noteSlot struct {
	note   types.Note
	ignore bool
}

type InMemoryNoteService struct {
	notes []noteSlot
}

func NewInMemoryNoteService() *InMemoryNoteService {
	return &InMemoryNoteService{
		notes: make([]noteSlot, 0),
	}
}

func (i *InMemoryNoteService) CreateNote(note types.Note) (int, error) {
	id := len(i.notes) + 1
	note.Id = id
	i.notes = append(i.notes, noteSlot{note, false})
	return id, nil
}

func (i *InMemoryNoteService) ViewNotes() []types.Note {
	notes := make([]types.Note, len(i.notes))
	for i, slot := range i.notes {
		notes[i] = slot.note
	}
	return notes
}

func (i *InMemoryNoteService) ViewSingleNote(id int) (types.Note, error) {
	if id > len(i.notes) {
		return types.Note{}, errors.New("No note found with that ID")
	}
	note := i.notes[id-1]
	if note.ignore {
		return types.Note{}, errors.New("Could not find note")
	}
	return note.note, nil
}

func (i *InMemoryNoteService) Delete(id int) error {
	if id > len(i.notes) {
		return nil
	}
	i.notes[id-1].ignore = true
	return nil
}
