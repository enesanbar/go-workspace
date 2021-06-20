package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/enesanbar/workspace/golang/projects/vigilate/config"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"
)

type HandlerLoginPost struct {
	Repository Repository
	Session    *session.Session
	Renderer   *services.Renderer
	cfg        *config.Config
}

func NewHandlerLoginPost(repository Repository, session *session.Session, renderer *services.Renderer, cfg *config.Config) *HandlerLoginPost {
	return &HandlerLoginPost{Repository: repository, Session: session, Renderer: renderer, cfg: cfg}
}

// ServeHTTP attempts to log the user in
func (h *HandlerLoginPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = h.Session.Manager.RenewToken(r.Context())
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		h.Renderer.ClientError(w, r, http.StatusBadRequest)
		return
	}

	id, hash, err := h.Repository.Authenticate(r.Form.Get("email"), r.Form.Get("password"))
	if err == models.ErrInvalidCredentials {
		h.Session.Manager.Put(r.Context(), "error", "Invalid login")
		err := h.Renderer.RenderPage(w, r, "login", nil, nil)
		if err != nil {
			h.Renderer.PrintTemplateError(w, err)
		}
		return
	} else if err == models.ErrInactiveAccount {
		h.Session.Manager.Put(r.Context(), "error", "Inactive account!")
		err := h.Renderer.RenderPage(w, r, "login", nil, nil)
		if err != nil {
			h.Renderer.PrintTemplateError(w, err)
		}
		return
	} else if err != nil {
		log.Println(err)
		h.Renderer.ClientError(w, r, http.StatusBadRequest)
		return
	}

	if r.Form.Get("remember") == "remember" {
		randomString := helpers.RandomString(12)
		hasher := sha256.New()

		_, err = hasher.Write([]byte(randomString))
		if err != nil {
			log.Println(err)
		}

		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

		err = h.Repository.InsertRememberMeToken(id, sha)
		if err != nil {
			log.Println(err)
		}

		// write a cookie
		expire := time.Now().Add(365 * 24 * 60 * 60 * time.Second)
		cookie := http.Cookie{
			Name:     fmt.Sprintf("_%s_gowatcher_remember", h.cfg.Identifier),
			Value:    fmt.Sprintf("%d|%s", id, sha),
			Path:     "/",
			Expires:  expire,
			HttpOnly: true,
			Domain:   h.cfg.Domain,
			MaxAge:   315360000, // seconds in year
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, &cookie)
	}

	// we authenticated. Get the user.
	u, err := h.Repository.GetUserById(id)
	if err != nil {
		log.Println(err)
		h.Renderer.ClientError(w, r, http.StatusBadRequest)
		return
	}

	h.Session.Manager.Put(r.Context(), "userID", id)
	h.Session.Manager.Put(r.Context(), "hashedPassword", hash)
	h.Session.Manager.Put(r.Context(), "flash", "You've been logged in successfully!")
	h.Session.Manager.Put(r.Context(), "user", u)

	if r.Form.Get("target") != "" {
		http.Redirect(w, r, r.Form.Get("target"), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/admin/overview", http.StatusSeeOther)
}
