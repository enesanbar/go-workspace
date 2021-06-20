package hosts

import (
	"log"
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

// ServeHTTP displays list of all hosts
func (h *HandlerAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hosts, err := h.Repository.GetHosts()
	if err != nil {
		log.Println(err)
		return
	}

	vars := make(jet.VarMap)
	vars.Set("hosts", hosts)
	err = h.Renderer.RenderPage(w, r, "hosts", vars, nil)
	if err != nil {
		h.Renderer.PrintTemplateError(w, err)
	}
}
