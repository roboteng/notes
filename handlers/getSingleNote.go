package handlers

import (
	"encoding/json"
	"net/http"
	"notes/types"

	"github.com/julienschmidt/httprouter"
)

func GetSingleNote(service types.SingleNoteViewer) func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		note, _ := service.ViewSingleNote(1)
		json_, _ := json.Marshal(note)
		w.WriteHeader(http.StatusOK)
		w.Write(json_)
	}
}
