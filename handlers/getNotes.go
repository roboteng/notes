package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notes/features"

	"github.com/julienschmidt/httprouter"
)

func GetNtes(feature *features.ViewNotes) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
