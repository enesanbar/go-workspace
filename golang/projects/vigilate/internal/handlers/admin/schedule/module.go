package schedule

import (
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	bindings,
)

var factories = fx.Provide(
	NewHandler,
)

var bindings = fx.Provide(
	func(repo repository.DatabaseRepo) Repository { return repo },
)
