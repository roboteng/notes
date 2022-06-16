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

func (i *InMemoryNoteService) Save(note types.Note) (int, error) {
	id := len(i.notes) + 1
	note.Id = id
	i.notes = append(i.notes, noteSlot{note, false})
	return id, nil
}

func (i *InMemoryNoteService) View() []types.Note {
	notes := make([]types.Note, len(i.notes))
	for i, slot := range i.notes {
		notes[i] = slot.note
	}
	return notes
}

func (i *InMemoryNoteService) ViewSingle(id int) (types.Note, error) {
	note, err := i.getNote(id)
	if err != nil {
		return types.Note{}, errors.New("No note found with that ID")
	}
	if note.ignore {
		return types.Note{}, errors.New("Could not find note")
	}
	return note.note, nil
}

func (i *InMemoryNoteService) Delete(id int) error {
	note, err := i.getNote(id)
	if err != nil {
		return nil
	}
	note.ignore = true
	return nil
}

func (i *InMemoryNoteService) Update(id int, newValues types.Note) error {
	note, err := i.getNote(id)
	if err != nil {
		return err
	}
	note.note = newValues
	return nil
}

func (i *InMemoryNoteService) getNote(id int) (*noteSlot, error) {
	if id > len(i.notes) {
		return nil, errors.New("id out of bounds")
	}
	return &i.notes[id-1], nil
}
