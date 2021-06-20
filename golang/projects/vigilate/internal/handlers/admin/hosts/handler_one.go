package hosts

import (
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

// ServeHTTP shows the host add/edit form
func (ho *HandlerOne) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var host models.Host
	if id > 0 {
		h, err := ho.Repository.GetHostByID(id)
		if err != nil {
			ho.Renderer.ServerError(w, r, err)
			return
		}
		host = h
	}

	vars := make(jet.VarMap)
	vars.Set("host", host)

	err := ho.Renderer.RenderPage(w, r, "host", vars, nil)
	if err != nil {
		ho.Renderer.PrintTemplateError(w, err)
	}
}
