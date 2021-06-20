package service

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
)

type HandlerProblemPage struct {
	repo     Repository
	renderer *services.Renderer
}

func NewHandlerProblemPage(repo Repository, renderer *services.Renderer) *HandlerProblemPage {
	return &HandlerProblemPage{repo: repo, renderer: renderer}
}

// AllProblemServices lists all warning services
func (h *HandlerProblemPage) AllProblemServices(w http.ResponseWriter, r *http.Request) {
	all, err := h.repo.GetServicesByStatus("problem")
	if err != nil {
		log.Println(err)
		return
	}

	vars := make(jet.VarMap)
	vars.Set("services", all)

	err = h.renderer.RenderPage(w, r, "problems", vars, nil)
	if err != nil {
		h.renderer.PrintTemplateError(w, err)
	}
}
