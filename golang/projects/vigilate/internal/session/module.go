package session

import (
	"github.com/alexedwards/scs/v2"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	scs.New,
	NewSessionStore,
	NewSession,
)
