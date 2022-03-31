package handlers

import (
	"fmt"
	"net/http"
	ty "notes/types"

	"github.com/julienschmidt/httprouter"
)

func CreateNote(service ty.NoteCreator) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		title := r.URL.Query().Get("title")
		if title != "" {
			id, err := service.CreateNote(ty.Note{})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "{\"id\":%d}", id)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}
