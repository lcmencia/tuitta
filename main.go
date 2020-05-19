package main

import (
	"log"

	"github.com/lcmencia/tuitta/bd"
	"github.com/lcmencia/tuitta/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("No connection to the database")
		return
	}
	handlers.Handlers()
}
