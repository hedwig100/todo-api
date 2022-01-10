package main

import (
	"log"

	"github.com/hedwig100/todo-app/cmd/app"
)

func main() {
	server := app.GetServer()
	log.Fatal(server.ListenAndServe())
}
