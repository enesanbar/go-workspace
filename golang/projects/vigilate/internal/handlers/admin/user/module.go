package user

import (
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	bindings,
)

var factories = fx.Provide(
	NewHandlerAll,
	NewHandlerOne,
	NewHandlerPost,
	NewHandlerDelete,
)

var bindings = fx.Provide(
	func(repo repository.DatabaseRepo) Repository { return repo },
)
