package main

import (
	"log"
)

func danger(args ...interface{}) {
	log.SetPrefix("ERROR ")
	log.Println(args...)
}
