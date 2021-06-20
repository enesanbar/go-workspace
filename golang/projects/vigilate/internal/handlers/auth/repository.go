package auth

import "github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

//go:generate mockery --name=Repository
type Repository interface {
	Authenticate(email, testPassword string) (int, string, error)
	InsertRememberMeToken(id int, token string) error
	GetUserById(id int) (models.User, error)
	DeleteToken(token string) error
}
