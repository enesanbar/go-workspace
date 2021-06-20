package main

import (
	"context"
	"encoding/gob"
	"log"
	"os"
	"runtime"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/enesanbar/workspace/golang/projects/vigilate/config"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository/dbrepo"

	"go.uber.org/fx"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
)

const vigilateVersion = "1.0.0"
const maxWorkerPoolSize = 5

func init() {
	gob.Register(models.User{})
	_ = os.Setenv("TZ", "Europe/Istanbul")
}

// main is the application entry point
func main() {
	cfg := config.NewConfig()

	// print info
	log.Printf("******************************************")
	log.Printf("** %sVigilate%s v%s built in %s", "\033[31m", "\033[0m", vigilateVersion, runtime.Version())
	log.Printf("**----------------------------------------")
	log.Printf("** Running with %d Processors", runtime.NumCPU())
	log.Printf("** Running on %s", runtime.GOOS)
	log.Printf("******************************************")

	app := fx.New(
		fx.Invoke(run),
		fx.Supply(cfg),
		handlers.Module,
		services.Module,
		session.Module,
		middlewares,
		fx.Provide(
			repository.NewDBConnectionConfig,
			repository.NewDBConnection,
			dbrepo.NewPostgresRepo,
			func(repo repository.DatabaseRepo) helpers.PrefRepository {
				return repo
			},
			NewRoutes,
			helpers.NewPreferences,

			fx.Annotated{
				Group:  "runnables",
				Target: services.NewDispatcher,
			},
			fx.Annotated{
				Group:  "runnables",
				Target: services.NewRunnableMonitoring,
			},
			fx.Annotated{
				Group:  "runnables",
				Target: NewServer,
			},
		),
	)

	app.Run()
}

type params struct {
	fx.In

	Runnables []helpers.Runnable `group:"runnables"`
}

func run(lc fx.Lifecycle, params params) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			for _, server := range params.Runnables {
				if server == nil {
					continue
				}
				go func(server helpers.Runnable) {
					log.Println("starting", server)
					err := server.Start()
					if err != nil {
						log.Println(err)
					}
				}(server)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			for _, server := range params.Runnables {
				if server == nil {
					continue
				}
				err := server.Stop()
				if err != nil {
					log.Println(err)
				}
			}
			return nil
		},
	})
}
