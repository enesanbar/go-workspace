package user

import "github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

type Repository interface {
	InsertUser(u models.User) (int, error)
	AllUsers() ([]*models.User, error)
	GetUserById(id int) (models.User, error)
	UpdatePassword(id int, newPassword string) error
	UpdateUser(u models.User) error
	DeleteUser(id int) error
}
