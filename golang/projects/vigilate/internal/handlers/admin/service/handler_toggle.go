package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
)

type serviceJson struct {
	OK bool `json:"ok"`
}

type HandlerToggleHostService struct {
	repo      Repository
	wsClient  *services.PusherClient
	prefs     *helpers.Preferences
	scheduler *services.Scheduler
}

func NewHandlerToggleHostService(repo Repository, wsClient *services.PusherClient, prefs *helpers.Preferences, scheduler *services.Scheduler) *HandlerToggleHostService {
	return &HandlerToggleHostService{repo: repo, wsClient: wsClient, prefs: prefs, scheduler: scheduler}
}

func (h *HandlerToggleHostService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	resp := serviceJson{
		OK: true,
	}

	hostID, _ := strconv.Atoi(r.Form.Get("host_id"))
	serviceID, _ := strconv.Atoi(r.Form.Get("service_id"))
	active, _ := strconv.Atoi(r.Form.Get("active"))

	err = h.repo.UpdateHostServiceStatus(hostID, serviceID, active)
	if err != nil {
		log.Println(err)
		resp.OK = false
		return
	}

	// broadcast
	hs, _ := h.repo.GetHostServiceByHostIDServiceID(hostID, serviceID)
	host, _ := h.repo.GetHostByID(hostID)

	// add or remove host service from schedule
	if active == 1 {
		h.wsClient.PushScheduleChangedEvent(hs, "pending")
		h.wsClient.PushStatusChangedEvent(host, hs, "pending")
		h.addToMonitorMap(hs)
	} else {
		h.removeFromMonitorMap(hs)
	}

	marshal, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}

func (h *HandlerToggleHostService) addToMonitorMap(hs models.HostService) {
	if h.prefs.GetPref("monitoring_live") == "0" {
		return
	}

	var j services.Job
	j.HostServiceID = hs.ID
	scheduleID, err := h.scheduler.Cron.AddJob(fmt.Sprintf("@every %d%s", hs.ScheduleNumber, hs.ScheduleUnit), j)
	if err != nil {
		log.Println(err)
		return
	}

	h.scheduler.MonitorMap[hs.ID] = scheduleID
	data := map[string]string{
		"message":         "scheduling",
		"host_service_id": strconv.Itoa(hs.ID),
		"next_run":        "Pending",
		"service":         hs.Service.ServiceName,
		"host":            hs.HostName,
		"last_run":        hs.LastCheck.Format("2006-01-02 3:04:05 PM"),
		"schedule":        fmt.Sprintf("@every %d%s", hs.ScheduleNumber, hs.ScheduleUnit),
	}

	h.wsClient.BroadcastMessage("public-channel", "schedule-changed-event", data)
}

func (h *HandlerToggleHostService) removeFromMonitorMap(hs models.HostService) {
	if h.prefs.GetPref("monitoring_live") == "0" {
		return
	}

	h.scheduler.Cron.Remove(h.scheduler.MonitorMap[hs.ID])
	data := map[string]string{
		"host_service_id": strconv.Itoa(hs.ID),
	}

	h.wsClient.BroadcastMessage("public-channel", "schedule-item-removed-event", data)
}
