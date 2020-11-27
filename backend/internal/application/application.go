package application

import "gorm.io/gorm"

// app struct contains global state.
type App struct {
	// db is the global database connection pool.
	db *gorm.DB
}

func InitApp(db *gorm.DB) *App {
	return &App{db}
}
