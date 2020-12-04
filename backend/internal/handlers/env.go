package handlers

import (
	"feedback/internal/auth"

	"gorm.io/gorm"
)

type Env struct {
	Db   *gorm.DB
	Auth auth.Authentication
}
