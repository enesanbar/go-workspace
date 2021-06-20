//+build wireinject

package main

import (
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/modules/exchange"
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/modules/get"
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/modules/list"
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/modules/register"
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/rest"
	"github.com/google/wire"
)

func initializeServer() (*rest.Server, error) {
	wire.Build(wireSet)
	return nil, nil
}

func initializeServerCustomConfig(_ exchange.Config, _ get.Config, _ list.Config, _ register.Config, _ rest.Config) *rest.Server {
	wire.Build(wireSetWithoutConfig)
	return nil
}
