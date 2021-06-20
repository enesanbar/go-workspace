package service

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
)

type HandlerPendingPage struct {
	repo     Repository
	renderer *services.Renderer
}

func NewHandlerPendingPage(repo Repository, renderer *services.Renderer) *HandlerPendingPage {
	return &HandlerPendingPage{repo: repo, renderer: renderer}
}

// AllPendingServices lists all warning services
func (h *HandlerPendingPage) AllPendingServices(w http.ResponseWriter, r *http.Request) {
	all, err := h.repo.GetServicesByStatus("pending")
	if err != nil {
		log.Println(err)
		return
	}

	vars := make(jet.VarMap)
	vars.Set("services", all)

	err = h.renderer.RenderPage(w, r, "pending", vars, nil)
	if err != nil {
		h.renderer.PrintTemplateError(w, err)
	}
}
