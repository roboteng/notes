package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "index.html")
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
