package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/enesanbar/workspace/golang/projects/movies/backend/models"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		host     string
		port     int
		database string
		user     string
		password string
	}
	jwt struct {
		secret string
	}
}

type AppStatus struct {
	Status      string `json:"status,omitempty"`
	Environment string `json:"environment,omitempty"`
	Version     string `json:"version,omitempty"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&cfg.db.host, "host", "localhost", "Database host")
	flag.IntVar(&cfg.db.port, "dbPort", 5432, "Database port")
	flag.StringVar(&cfg.db.user, "dbUser", "postgres", "Database user")
	flag.StringVar(&cfg.db.password, "password", "", "Database password")
	flag.StringVar(&cfg.db.database, "database", "movies", "Database name")
	flag.Parse()

	cfg.jwt.secret = os.Getenv("GO_MOVIES_JWT")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := openDB(cfg)
	if err != nil {
		logger.Println(err)
		return
	}
	defer db.Close()

	app := application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Println("Starting server on port ", cfg.port)

	if err := srv.ListenAndServe(); err != nil {
		app.logger.Println(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.db.user, cfg.db.password, cfg.db.host, cfg.db.port, cfg.db.database,
		),
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, err
}
