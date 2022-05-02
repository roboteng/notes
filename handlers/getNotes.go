package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notes/features"
	ty "notes/types"

	"github.com/julienschmidt/httprouter"
)

func GetNtes(service ty.NotesViewer) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	feature := features.ViewNotes{Service: service}
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		notes, err := feature.View()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		bytes, _ := json.Marshal(notes)
		fmt.Fprint(w, string(bytes))
	}
}
