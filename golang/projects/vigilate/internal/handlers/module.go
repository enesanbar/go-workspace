package handlers

import (
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/dashboard"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/events"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/hosts"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/prefs"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/pusher"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/schedule"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/service"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/settings"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/user"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/auth"
	"go.uber.org/fx"
)

var Module = fx.Options(
	dashboard.Module,
	events.Module,
	hosts.Module,
	prefs.Module,
	pusher.Module,
	schedule.Module,
	service.Module,
	settings.Module,
	user.Module,
	auth.Module,
)
