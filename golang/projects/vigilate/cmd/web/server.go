package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/enesanbar/workspace/golang/projects/vigilate/config"
)

type ServerConfig struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
}

type Server struct {
	server *http.Server
	cfg    *config.Config
}

func NewServer(cfg *config.Config, routes http.Handler) helpers.Runnable {
	// create http server
	srv := &http.Server{
		Addr:              cfg.Port,
		Handler:           routes,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	return &Server{
		server: srv,
		cfg:    cfg,
	}
}

func (s *Server) Start() error {
	log.Printf("Starting HTTP server on port %s....", s.cfg.Port)

	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	log.Println("Stopping HTTP Server")
	if err := s.server.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("error shutting down Server (%w)", err)
	}
	return nil
}
