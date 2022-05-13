package handlers

import (
	"net/http"
	"notes/features"
	ty "notes/types"

	"github.com/julienschmidt/httprouter"
)

func MakeRouter(service ty.NoteService) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.ServeFile(w, r, "index.html")
	})
	router.GET("/api/notes", GetNtes(&features.ViewNotes{Service: service}))
	router.POST("/api/notes", CreateNote(service))
	router.GET("/api/notes/:id", GetSingleNote(service))
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	return router
}
