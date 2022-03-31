package main

import (
	"log"
	"net/http"
	"notes/handlers"
	"notes/types"
)

type InMemoryNoteService struct {
	notes []types.Note
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

func main() {
	service := &InMemoryNoteService{}
	router := handlers.MakeRouter(service)
	log.Fatal(http.ListenAndServe(":8080", router))
}
