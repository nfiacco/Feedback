package application

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

// app struct contains global state.
type App struct {
	// db is the global database connection pool.
	db *gorm.DB
}

func InitApp(db *gorm.DB) *App {
	rand.Seed(time.Now().UTC().UnixNano())

	return &App{db}
}
