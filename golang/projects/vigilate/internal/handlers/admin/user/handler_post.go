package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"
	"github.com/go-chi/chi/v5"
)

type HandlerPost struct {
	Repository Repository
	Renderer   *services.Renderer
	Session    *session.Session
}

func NewHandlerPost(repository Repository, renderer *services.Renderer, session *session.Session) *HandlerPost {
	return &HandlerPost{Repository: repository, Renderer: renderer, Session: session}
}

// ServeHTTP adds/edits a user
func (h *HandlerPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
	}

	var u models.User

	if id > 0 {
		u, _ = h.Repository.GetUserById(id)
		u.FirstName = r.Form.Get("first_name")
		u.LastName = r.Form.Get("last_name")
		u.Email = r.Form.Get("email")
		u.UserActive, _ = strconv.Atoi(r.Form.Get("user_active"))
		err := h.Repository.UpdateUser(u)
		if err != nil {
			log.Println(err)
			h.Renderer.ClientError(w, r, http.StatusBadRequest)
			return
		}

		if len(r.Form.Get("password")) > 0 {
			// changing password
			err := h.Repository.UpdatePassword(id, r.Form.Get("password"))
			if err != nil {
				log.Println(err)
				h.Renderer.ClientError(w, r, http.StatusBadRequest)
				return
			}
		}
	} else {
		u.FirstName = r.Form.Get("first_name")
		u.LastName = r.Form.Get("last_name")
		u.Email = r.Form.Get("email")
		u.UserActive, _ = strconv.Atoi(r.Form.Get("user_active"))
		u.Password = []byte(r.Form.Get("password"))
		u.AccessLevel = 3

		_, err := h.Repository.InsertUser(u)
		if err != nil {
			log.Println(err)
			h.Renderer.ClientError(w, r, http.StatusBadRequest)
			return
		}
	}

	h.Session.Manager.Put(r.Context(), "flash", "Changes saved")
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
