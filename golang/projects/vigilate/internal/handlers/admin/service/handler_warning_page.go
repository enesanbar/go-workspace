package service

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
)

type HandlerWarningPage struct {
	repo     Repository
	renderer *services.Renderer
}

func NewHandlerWarningPage(repo Repository, renderer *services.Renderer) *HandlerWarningPage {
	return &HandlerWarningPage{repo: repo, renderer: renderer}
}

// AllWarningServices lists all warning services
func (h *HandlerWarningPage) AllWarningServices(w http.ResponseWriter, r *http.Request) {
	all, err := h.repo.GetServicesByStatus("warning")
	if err != nil {
		log.Println(err)
		return
	}

	vars := make(jet.VarMap)
	vars.Set("services", all)

	err = h.renderer.RenderPage(w, r, "warning", vars, nil)
	if err != nil {
		h.renderer.PrintTemplateError(w, err)
	}
}
