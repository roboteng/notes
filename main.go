package main

import (
	"log"
	"net/http"
	"notes/handlers"
	"notes/services"
)

func main() {
	service := &services.InMemoryNoteService{}
	router := handlers.MakeRouter(service)
	log.Fatal(http.ListenAndServe(":8080", router))
}
