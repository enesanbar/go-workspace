package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/enesanbar/workspace/golang/projects/vigilate/pkg/driver"

	"github.com/enesanbar/workspace/golang/projects/vigilate/config"
)

type DBConnection struct {
	Conn *sql.DB
}

type DBConnectionConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	DBSSL  string
}

func NewDBConnectionConfig(cfg *config.Config) *DBConnectionConfig {
	return &DBConnectionConfig{
		DBHost: cfg.DBConfig.Host,
		DBPort: cfg.DBConfig.Port,
		DBUser: cfg.DBConfig.User,
		DBPass: cfg.DBConfig.Pass,
		DBName: cfg.DBConfig.Name,
		DBSSL:  cfg.DBConfig.SSL,
	}
}

func NewDBConnection(cfg *DBConnectionConfig) *DBConnection {
	log.Println("Connecting to database....")
	dsnString := ""

	// when developing locally, we often don't have a db password
	if cfg.DBPass == "" {
		dsnString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBName,
			cfg.DBSSL,
		)
	} else {
		dsnString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBPass,
			cfg.DBName,
			cfg.DBSSL,
		)
	}

	db, err := driver.ConnectPostgres(dsnString)
	if err != nil {
		log.Fatal("Cannot connect to database!", err)
	}

	return &DBConnection{Conn: db}
}
