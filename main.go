package main

import (
	"log"
	"net/http"
	"notes/handlers"
)

func main() {
	router := handlers.MakeRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
