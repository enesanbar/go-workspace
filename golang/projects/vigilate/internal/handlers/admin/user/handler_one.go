package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/CloudyKit/jet/v6"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/go-chi/chi/v5"
)

type HandlerOne struct {
	Repository Repository
	Renderer   *services.Renderer
}

func NewHandlerOne(repository Repository, renderer *services.Renderer) *HandlerOne {
	return &HandlerOne{Repository: repository, Renderer: renderer}
}

// ServeHTTP displays the add/edit user page
func (h *HandlerOne) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
	}

	vars := make(jet.VarMap)

	if id > 0 {

		u, err := h.Repository.GetUserById(id)
		if err != nil {
			h.Renderer.ClientError(w, r, http.StatusBadRequest)
			return
		}

		vars.Set("user", u)
	} else {
		var u models.User
		vars.Set("user", u)
	}

	err = h.Renderer.RenderPage(w, r, "user", vars, nil)
	if err != nil {
		h.Renderer.PrintTemplateError(w, err)
	}
}
