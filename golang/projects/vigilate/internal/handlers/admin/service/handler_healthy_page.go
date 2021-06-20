package service

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
)

type HandlerHealthyPage struct {
	repo     Repository
	renderer *services.Renderer
}

func NewHandlerHealthyPage(repo Repository, renderer *services.Renderer) *HandlerHealthyPage {
	return &HandlerHealthyPage{repo: repo, renderer: renderer}
}

// AllHealthyServices lists all healthy services
func (h *HandlerHealthyPage) AllHealthyServices(w http.ResponseWriter, r *http.Request) {
	all, err := h.repo.GetServicesByStatus("healthy")
	if err != nil {
		log.Println(err)
		return
	}

	vars := make(jet.VarMap)
	vars.Set("services", all)

	err = h.renderer.RenderPage(w, r, "healthy", vars, nil)
	if err != nil {
		h.renderer.PrintTemplateError(w, err)
	}
}
