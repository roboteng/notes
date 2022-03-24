package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Note struct {
	Title    string `json:"title"`
	Contents string `json:"desc"`
}

func main() {
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "index.html")
	}))
	http.Handle("/api/notes", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notes := []Note{
			{
				"Note 1",
				"Some info",
			},
			{
				"Note 2",
				"Some things",
			},
		}
		bytes, _ := json.Marshal(notes)
		fmt.Fprint(w, string(bytes))
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
