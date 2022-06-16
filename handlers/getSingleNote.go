package handlers

import (
	"encoding/json"
	"net/http"
	"notes/types"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetSingleNote(service types.SingleNoteViewer) func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id_ := params.ByName("id")
		id, _ := strconv.Atoi(id_)
		note, err := service.ViewSingle(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json_, _ := json.Marshal(note)
		w.WriteHeader(http.StatusOK)
		w.Write(json_)
	}
}
