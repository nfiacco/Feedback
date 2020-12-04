package main

import (
	"feedback/internal/database"
	"feedback/internal/router"
	"log"
	"math/rand"
	"time"
)

func main() {

	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())

	router.RunServer(db)
}
