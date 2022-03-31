package main

import (
	"log"
	"net/http"
	"notes/handlers"
	"notes/types"
)

func main() {
	service := &types.AnonNoteService{
		AnonNoteCreator: types.AnonNoteCreator{
			Create: func(note types.Note) (int, error) {
				panic("Not Implemented")
			}},
		AnonNotesViewer: types.AnonNotesViewer{
			View: func() []types.Note {
				return make([]types.Note, 0)
			},
		},
	}
	router := handlers.MakeRouter(service)
	log.Fatal(http.ListenAndServe(":8080", router))
}
