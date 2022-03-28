package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func MakeRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.ServeFile(w, r, "index.html")
	})
	router.GET("/api/notes", GetNtes())

	return router
}
