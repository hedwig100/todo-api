package main

import (
	"log"
	"net/http"

	"github.com/hedwig100/todo-app/internal/route"
)

func main() {
	mux := route.GetMux()
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &mux,
	}
	log.Fatal(server.ListenAndServe())
}
