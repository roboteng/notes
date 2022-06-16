package services_test

import (
	. "notes/services"
	"notes/types"
	"testing"
)

func TestLocalStore(t *testing.T) {
	t.Run("After a note is created, that same ID should match the stored note", func(t *testing.T) {
		store := NewInMemoryNoteService()
		note := types.Note{Title: "First Note", Contents: "I made a note!"}
		id, err := store.CreateNote(note)
		if err != nil {
			t.Error(err)
		}
		gotNote, err := store.ViewSingleNote(id)
		if err != nil {
			t.Error(err)
		}
		if note.Title != gotNote.Title && note.Contents != note.Contents {
			t.Error("Got the incorrect note")
		}
	})

	t.Run("Returns an error if asked for a note that doesn't exist", func(t *testing.T) {
		store := NewInMemoryNoteService()
		_, err := store.ViewSingleNote(1)
		if err != nil {
			return
		}
		t.Error("Expected an error when asking for a note that doesn't exist")
	})
}
