package events

import (
	"log"
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/CloudyKit/jet/v6"
)

type HandlerGet struct {
	repo     Repository
	renderer *services.Renderer
}

func NewHandlerGet(repo Repository, renderer *services.Renderer) *HandlerGet {
	return &HandlerGet{repo: repo, renderer: renderer}
}

// Events displays the events page
func (h *HandlerGet) Events(w http.ResponseWriter, r *http.Request) {
	events, err := h.repo.GetAllEvents()
	if err != nil {
		log.Println(err)
		return
	}

	vars := make(jet.VarMap)
	vars.Set("events", events)

	err = h.renderer.RenderPage(w, r, "events", vars, nil)
	if err != nil {
		h.renderer.PrintTemplateError(w, err)
	}
}
