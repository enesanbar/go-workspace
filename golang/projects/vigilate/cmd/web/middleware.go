package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.uber.org/fx"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/enesanbar/workspace/golang/projects/vigilate/config"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"
	"github.com/justinas/nosurf"
)

var middlewares = fx.Options(
	fx.Provide(
		fx.Annotated{
			Group:  "middlewares",
			Target: NewRecoverPanic,
		},
		fx.Annotated{
			Group:  "middlewares",
			Target: NewNoSurf,
		},
		fx.Annotated{
			Group:  "middlewares",
			Target: NewCheckRemember,
		},
	),
)

// NewAuth checks for authentication
func NewAuth(session *session.Session) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !session.IsAuthenticated(r) {
				url := r.URL.Path
				http.Redirect(w, r, fmt.Sprintf("/?target=%s", url), http.StatusFound)
				return
			}
			w.Header().Add("Cache-Control", "no-store")

			next.ServeHTTP(w, r)
		})
	}
}

func NewRecoverPanic(renderer *services.Renderer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				// Check if there has been a panic
				if err := recover(); err != nil {
					// return a 500 Internal Server response
					renderer.ServerError(w, r, fmt.Errorf("%s", err))
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// NewNoSurf implements CSRF protection
func NewNoSurf(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		csrfHandler := nosurf.New(next)

		csrfHandler.ExemptPath("/pusher/auth")
		csrfHandler.ExemptPath("/pusher/hook")

		csrfHandler.SetBaseCookie(http.Cookie{
			HttpOnly: true,
			Path:     "/",
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
			Domain:   cfg.Domain,
		})

		return csrfHandler
	}
}

// NewCheckRemember checks to see if we should log the user in automatically
func NewCheckRemember(session *session.Session, repo repository.DatabaseRepo, prefs *helpers.Preferences, cfg *config.Config) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !session.IsAuthenticated(r) {
				cookie, err := r.Cookie(fmt.Sprintf("_%s_gowatcher_remember", cfg.Identifier))
				if err != nil {
					next.ServeHTTP(w, r)
				} else {
					key := cookie.Value
					// have a remember token, so try to log the user in
					if len(key) > 0 {
						// key length > 0, so it might br a valid token
						split := strings.Split(key, "|")
						uid, hash := split[0], split[1]
						id, _ := strconv.Atoi(uid)
						validHash := repo.CheckForToken(id, hash)
						if validHash {
							// valid remember me token, so log the user in
							_ = session.Manager.RenewToken(r.Context())
							user, _ := repo.GetUserById(id)
							hashedPassword := user.Password
							session.Manager.Put(r.Context(), "userID", id)
							session.Manager.Put(r.Context(), "userName", user.FirstName)
							session.Manager.Put(r.Context(), "userFirstName", user.FirstName)
							session.Manager.Put(r.Context(), "userLastName", user.LastName)
							session.Manager.Put(r.Context(), "hashedPassword", string(hashedPassword))
							session.Manager.Put(r.Context(), "user", user)
							next.ServeHTTP(w, r)
						} else {
							// invalid token, so delete the cookie
							deleteRememberCookie(w, r, session, prefs, cfg)
							session.Manager.Put(r.Context(), "error", "You've been logged out from another device!")
							next.ServeHTTP(w, r)
						}
					} else {
						// key length is zero, so it's a leftover cookie (user has not closed browser)
						next.ServeHTTP(w, r)
					}
				}
			} else {
				// they are logged in, but make sure that the remember token has not been revoked
				cookie, err := r.Cookie(fmt.Sprintf("_%s_gowatcher_remember", cfg.Identifier))
				if err != nil {
					// no cookie
					next.ServeHTTP(w, r)
				} else {
					key := cookie.Value
					// have a remember token, but make sure it's valid
					if len(key) > 0 {
						split := strings.Split(key, "|")
						uid, hash := split[0], split[1]
						id, _ := strconv.Atoi(uid)
						validHash := repo.CheckForToken(id, hash)
						if !validHash {
							deleteRememberCookie(w, r, session, prefs, nil)
							session.Manager.Put(r.Context(), "error", "You've been logged out from another device!")
							next.ServeHTTP(w, r)
						} else {
							next.ServeHTTP(w, r)
						}
					} else {
						next.ServeHTTP(w, r)
					}
				}
			}
		})
	}
}

// deleteRememberCookie deletes the remember me cookie, and logs the user out
func deleteRememberCookie(w http.ResponseWriter, r *http.Request, session *session.Session, prefs *helpers.Preferences, cfg *config.Config) {
	_ = session.Manager.RenewToken(r.Context())
	// delete the cookie
	newCookie := http.Cookie{
		Name:     fmt.Sprintf("_%s_ggowatcher_remember", cfg.Identifier),
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-100 * time.Hour),
		HttpOnly: true,
		Domain:   cfg.Domain,
		MaxAge:   -1,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &newCookie)

	// log them out
	session.Manager.Remove(r.Context(), "userID")
	_ = session.Manager.Destroy(r.Context())
	_ = session.Manager.RenewToken(r.Context())
}
