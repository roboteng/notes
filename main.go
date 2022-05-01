package main

import (
	"fmt"
	"log"
	"net/http"
	"notes/handlers"
	"notes/services"
)

func main() {
	service := services.NewInMemoryNoteService()
	router := handlers.MakeRouter(service)
	fmt.Println("Serving on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
