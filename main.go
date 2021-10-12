package main

import (
	"log"
)

func main() {
	server := server()
	log.Fatal("[Fatal] [main]", server.ListenAndServe())
}
