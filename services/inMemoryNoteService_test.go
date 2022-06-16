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
		id, err := store.Save(note)
		if err != nil {
			t.Error(err)
		}
		gotNote, err := store.ViewSingle(id)
		if err != nil {
			t.Error(err)
		}
		if !note.Equals(&gotNote) {
			t.Error("Got the incorrect note")
		}
	})

	t.Run("Returns an error if asked for a note that doesn't exist", func(t *testing.T) {
		store := NewInMemoryNoteService()
		_, err := store.ViewSingle(1)
		if err != nil {
			return
		}
		t.Error("Expected an error when asking for a note that doesn't exist")
	})

	t.Run("Deleting a non-existant note does not return an error", func(t *testing.T) {
		store := NewInMemoryNoteService()
		err := store.Delete(1)
		if err != nil {
			t.Error("Did no expect an error when deleting an invalid id")
		}
	})

	t.Run("Getting an ID for a note that has been deleted returns an error", func(t *testing.T) {
		store := NewInMemoryNoteService()
		note := types.Note{Title: "Sample Title", Contents: "I made a note"}
		id, _ := store.Save(note)
		store.Delete(id)
		_, err := store.ViewSingle(id)
		if err != nil {
			return
		}
		t.Error("Expected an error when looking for a deleted note")
	})

	t.Run("When two notes are stored, after deleting the first, the second still exists", func(t *testing.T) {
		store := NewInMemoryNoteService()
		idToDelete, _ := store.Save(types.Note{Title: "", Contents: ""})
		noteToKeep := types.Note{Title: "Second Note", Contents: "I made a note"}
		idToKeep, _ := store.Save(noteToKeep)
		store.Delete(idToDelete)
		note, err := store.ViewSingle(idToKeep)
		if err != nil {
			t.Error("Expected to find a note, even though the other was deleted")
		}
		if !note.Equals(&noteToKeep) {
			t.Errorf("Got %v, but expected %v", note, noteToKeep)
		}
	})

	t.Run("When updating an note, the note should save the update", func(t *testing.T) {
		store := NewInMemoryNoteService()
		id, _ := store.Save(types.Note{Title: "Old Title", Contents: "Old Contents"})
		updatedNote := types.Note{Title: "New Title", Contents: "New Contents"}
		err := store.Update(id, updatedNote)
		if err != nil {
			t.Error("Did not expect an error when updating the note")
		}
		note, _ := store.ViewSingle(id)
		if !note.Equals(&updatedNote) {
			t.Errorf("Got %v, but wanted %v", note, updatedNote)
		}
	})
}
