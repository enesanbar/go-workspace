package dbrepo

import (
	"database/sql"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"
)

type postgresDBRepo struct {
	DB *sql.DB
}

// NewPostgresRepo creates the repository
func NewPostgresRepo(conn *repository.DBConnection) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn.Conn,
	}
}
