package pusher

import "github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

type Repository interface {
	GetUserById(id int) (models.User, error)
}
