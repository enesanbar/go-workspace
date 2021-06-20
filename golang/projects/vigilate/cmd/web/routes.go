package main

import (
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/prefs"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/events"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/service"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/pusher"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/schedule"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/settings"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/auth"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/user"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/hosts"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/handlers/admin/dashboard"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"

	"go.uber.org/fx"

	"github.com/go-chi/chi/v5"
)

type RouteParams struct {
	fx.In

	Session   *session.Session
	Dashboard *dashboard.Handler
	HostAll   *hosts.HandlerAll
	HostOne   *hosts.HandlerOne
	HostPost  *hosts.HandlerPost

	UserPost   *user.HandlerPost
	UserAll    *user.HandlerAll
	UserOne    *user.HandlerOne
	UserDelete *user.HandlerDelete

	SettingGet  *settings.HandlerGet
	SettingPost *settings.HandlerPost

	AuthLoginGet  *auth.HandlerLoginGet
	AuthLoginPost *auth.HandlerLoginPost
	AuthLogout    *auth.HandlerLogout

	ScheduleGet *schedule.Handler

	PrefsToggleMonitoring *prefs.HandlerToggleMonitoring
	SetPref               *prefs.HandlerSetPref

	PusherAuth    *pusher.HandlerAuth
	PusherPrivate *pusher.HandlerPrivate
	TestPusher    *pusher.HandlerTestPusher

	WarningPage       *service.HandlerWarningPage
	HealthyPage       *service.HandlerHealthyPage
	ProblemPage       *service.HandlerProblemPage
	PendingPage       *service.HandlerPendingPage
	ToggleHostService *service.HandlerToggleHostService
	TestCheckNow      *service.HandlerTestCheckNow

	EventPage *events.HandlerGet

	Middlewares []func(next http.Handler) http.Handler `group:"middlewares"`
}

func NewRoutes(p RouteParams) http.Handler {

	mux := chi.NewRouter()

	// default middleware
	for _, m := range p.Middlewares {
		mux.Use(m)
	}

	// login
	mux.Get("/", p.AuthLoginGet.ServeHTTP)
	mux.Post("/", p.AuthLoginPost.ServeHTTP)
	mux.Get("/user/logout", p.AuthLogout.ServeHTTP)

	mux.Get("/pusher/test", p.TestPusher.ServeHTTP)

	authMiddleware := NewAuth(p.Session)
	mux.Route("/pusher", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Post("/auth", p.PusherAuth.ServeHTTP)
	})

	// admin NewRoutes
	mux.Route("/admin", func(mux chi.Router) {
		// all admin NewRoutes are protected
		mux.Use(authMiddleware)

		// test code to send private message
		mux.Get("/private-message", p.PusherPrivate.SendPrivateMessage)

		// overview
		mux.Get("/overview", p.Dashboard.ServeHTTP)

		// events
		mux.Get("/events", p.EventPage.Events)

		// settings
		mux.Get("/settings", p.SettingGet.ServeHTTP)
		mux.Post("/settings", p.SettingPost.ServeHTTP)

		// service status pages (all hosts)
		mux.Get("/all-healthy", p.HealthyPage.AllHealthyServices)
		mux.Get("/all-warning", p.WarningPage.AllWarningServices)
		mux.Get("/all-problems", p.ProblemPage.AllProblemServices)
		mux.Get("/all-pending", p.PendingPage.AllPendingServices)

		// users
		mux.Get("/users", p.UserAll.ServeHTTP)
		mux.Get("/user/{id}", p.UserOne.ServeHTTP)
		mux.Post("/user/{id}", p.UserPost.ServeHTTP)
		mux.Get("/user/delete/{id}", p.UserDelete.ServeHTTP)

		// schedule
		mux.Get("/schedule", p.ScheduleGet.ServeHTTP)

		// preferences
		mux.Post("/preference/ajax/set-system-pref", p.SetPref.SetSystemPref)
		mux.Post("/preference/ajax/toggle-monitoring", p.PrefsToggleMonitoring.ToggleMonitoring)

		// hosts
		mux.Get("/host/all", p.HostAll.ServeHTTP)
		mux.Get("/host/{id}", p.HostOne.ServeHTTP)
		mux.Post("/host/{id}", p.HostPost.ServeHTTP)
		mux.Post("/host/ajax/toggle-service", p.ToggleHostService.ServeHTTP)
		mux.Get("/perform-check/{id}/{oldStatus}", p.TestCheckNow.ServeHTTP)
	})

	// static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return p.Session.Manager.LoadAndSave(mux)
}
