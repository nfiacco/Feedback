package main

import (
	"feedback/internal/config"
	"feedback/internal/database"
	"feedback/internal/router"
	"log"
	"math/rand"
	"time"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())

	router.RunServer(db)
}
