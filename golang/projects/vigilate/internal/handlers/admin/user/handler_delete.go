package user

import (
	"net/http"
	"strconv"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"
	"github.com/go-chi/chi/v5"
)

type HandlerDelete struct {
	Repository Repository
	Session    *session.Session
}

func NewHandlerDelete(repository Repository, session *session.Session) *HandlerDelete {
	return &HandlerDelete{Repository: repository, Session: session}
}

// ServeHTTP soft deletes a user
func (h *HandlerDelete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_ = h.Repository.DeleteUser(id)
	h.Session.Manager.Put(r.Context(), "flash", "User deleted")
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
