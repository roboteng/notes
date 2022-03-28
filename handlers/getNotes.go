package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	ty "notes/types"

	"github.com/julienschmidt/httprouter"
)

func GetNtes(service ty.NotesViewer) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		notes := service.ViewNotes()
		bytes, _ := json.Marshal(notes)
		fmt.Fprint(w, string(bytes))
	}
}
