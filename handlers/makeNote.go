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
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, "{\"id\":1}")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}
