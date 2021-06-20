package main

import (
	"context"
	"os"

	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/config"
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/modules/exchange"
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/modules/get"
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/modules/list"
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/modules/register"
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/rest"
	"github.com/google/wire"
)

// List of wire enabled objects
var wireSetWithoutConfig = wire.NewSet(
	// *exchange.Converter
	exchange.NewConverter,

	// *get.Getter
	get.NewGetter,

	// *list.Lister
	list.NewLister,

	// *register.Registerer
	wire.Bind(new(register.Exchanger), new(*exchange.Converter)),
	register.NewRegisterer,

	// *rest.Server
	wire.Bind(new(rest.GetModel), new(*get.Getter)),
	wire.Bind(new(rest.ListModel), new(*list.Lister)),
	wire.Bind(new(rest.RegisterModel), new(*register.Registerer)),
	rest.New,
)

var wireSet = wire.NewSet(
	wireSetWithoutConfig,

	// *config.Config
	config.Load,

	// *exchange.Converter
	wire.Bind(new(exchange.Config), new(*config.Config)),

	// *get.Getter
	wire.Bind(new(get.Config), new(*config.Config)),

	// *list.Lister
	wire.Bind(new(list.Config), new(*config.Config)),

	// *register.Registerer
	wire.Bind(new(register.Config), new(*config.Config)),

	// *rest.Server
	wire.Bind(new(rest.Config), new(*config.Config)),
)

func main() {
	// bind stop channel to context
	ctx := context.Background()

	// start REST server
	server, err := initializeServer()
	if err != nil {
		os.Exit(-1)
	}

	server.Listen(ctx.Done())
}
