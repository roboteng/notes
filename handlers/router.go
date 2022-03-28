package handlers

import (
	"net/http"
	ty "notes/types"

	"github.com/julienschmidt/httprouter"
)

func MakeRouter(service ty.NotesViewer) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.ServeFile(w, r, "index.html")
	})
	router.GET("/api/notes", GetNtes(service))

	return router
}
