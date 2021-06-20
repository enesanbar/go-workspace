package schedule

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/CloudyKit/jet/v6"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
)

type Handler struct {
	scheduler *services.Scheduler
	repo      Repository
	renderer  *services.Renderer
}

func NewHandler(scheduler *services.Scheduler, repo Repository, renderer *services.Renderer) *Handler {
	return &Handler{scheduler: scheduler, repo: repo, renderer: renderer}
}

// ServeHTTP lists schedule entries
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var items []models.Schedule
	for k, v := range h.scheduler.MonitorMap {
		hs, err := h.repo.GetHostServiceByID(k)
		if err != nil {
			log.Println(err)
			return
		}

		item := models.Schedule{
			ID:            k,
			EntryID:       v,
			Entry:         h.scheduler.Cron.Entry(v),
			ScheduleText:  fmt.Sprintf("@every %d%s", hs.ScheduleNumber, hs.ScheduleUnit),
			Host:          hs.HostName,
			Service:       hs.Service.ServiceName,
			LastRunFromHS: hs.LastCheck,
		}

		items = append(items, item)
	}

	sort.Sort(models.ScheduleByHost(items))
	vars := make(jet.VarMap)
	vars.Set("items", items)

	err := h.renderer.RenderPage(w, r, "schedule", vars, nil)
	if err != nil {
		h.renderer.PrintTemplateError(w, err)
	}
}
