package dashboard

import (
	"log"
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/CloudyKit/jet/v6"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"
)

type Repository interface {
	GetAllServiceStatusCounts() (int, int, int, int, error)
	GetHosts() ([]*models.Host, error)
}

type Handler struct {
	Repository Repository
	Renderer   *services.Renderer
}

func NewHandler(repository repository.DatabaseRepo, renderer *services.Renderer) *Handler {
	return &Handler{Repository: repository, Renderer: renderer}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pending, healthy, warning, problem, err := h.Repository.GetAllServiceStatusCounts()
	if err != nil {
		log.Println(err)
		return
	}

	vars := make(jet.VarMap)
	vars.Set("no_pending", pending)
	vars.Set("no_healthy", healthy)
	vars.Set("no_warning", warning)
	vars.Set("no_problem", problem)

	allHosts, err := h.Repository.GetHosts()
	if err != nil {
		log.Println(err)
		return
	}

	vars.Set("hosts", allHosts)

	err = h.Renderer.RenderPage(w, r, "dashboard", vars, nil)
	if err != nil {
		h.Renderer.PrintTemplateError(w, err)
	}
}
