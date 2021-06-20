package user

import (
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/CloudyKit/jet/v6"
)

type HandlerAll struct {
	Repository Repository
	Renderer   *services.Renderer
}

func NewHandlerAll(repository Repository, renderer *services.Renderer) *HandlerAll {
	return &HandlerAll{Repository: repository, Renderer: renderer}
}

// ServeHTTP lists all admin users
func (h *HandlerAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := make(jet.VarMap)

	u, err := h.Repository.AllUsers()
	if err != nil {
		h.Renderer.ClientError(w, r, http.StatusBadRequest)
		return
	}

	vars.Set("users", u)

	err = h.Renderer.RenderPage(w, r, "users", vars, nil)
	if err != nil {
		h.Renderer.PrintTemplateError(w, err)
	}
}
