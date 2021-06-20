package prefs

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
)

type HandlerToggleMonitoring struct {
	prefs      *helpers.Preferences
	monitoring *services.Monitoring
	scheduler  *services.Scheduler
	wsClient   *services.PusherClient
}

func NewHandlerToggleMonitoring(prefs *helpers.Preferences, monitoring *services.Monitoring, scheduler *services.Scheduler, wsClient *services.PusherClient) *HandlerToggleMonitoring {
	return &HandlerToggleMonitoring{prefs: prefs, monitoring: monitoring, scheduler: scheduler, wsClient: wsClient}
}

// ToggleMonitoring turns monitoring on and off
func (h *HandlerToggleMonitoring) ToggleMonitoring(w http.ResponseWriter, r *http.Request) {
	enabled := r.PostForm.Get("enabled")
	log.Println("Monitoring enabled:", enabled)

	if enabled == "1" {
		// start monitoring
		log.Println("turning monitoring on")
		h.prefs.SetPref("monitoring_live", "1")
		h.monitoring.Start()
		h.scheduler.Start()
	} else {
		// stop monitoring
		log.Println("turning monitoring off")
		h.prefs.SetPref("monitoring_live", "0")

		// remove all items in scheduler map
		for _, value := range h.scheduler.MonitorMap {
			h.scheduler.Cron.Remove(value)
		}

		for key := range h.scheduler.MonitorMap {
			delete(h.scheduler.MonitorMap, key)
		}

		for _, entry := range h.scheduler.Cron.Entries() {
			h.scheduler.Cron.Remove(entry.ID)
		}

		h.scheduler.Cron.Stop()
		// trigger a message to broadcast to all clients that app is starting to monitor
		data := make(map[string]string)
		data["message"] = "Monitoring is off..."
		err := h.wsClient.Client.Trigger("public-channel", "app-stopping", data)
		if err != nil {
			log.Println(err)
		}
	}

	resp := models.TestCheckResp{
		OK: true,
	}

	marshal, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}
