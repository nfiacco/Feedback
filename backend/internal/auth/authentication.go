package auth

import "feedback/internal/models"

type Authentication struct {
	Session         *models.Session
	IsAuthenticated bool
}
