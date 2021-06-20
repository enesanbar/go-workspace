package settings

import (
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
)

type HandlerGet struct {
	Renderer *services.Renderer
}

func NewHandlerGet(renderer *services.Renderer) *HandlerGet {
	return &HandlerGet{Renderer: renderer}
}

// ServeHTTP displays the settings page
func (h *HandlerGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Renderer.RenderPage(w, r, "settings", nil, nil)
	if err != nil {
		h.Renderer.PrintTemplateError(w, err)
	}
}
