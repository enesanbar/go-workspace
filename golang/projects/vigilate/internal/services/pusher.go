package services

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/enesanbar/workspace/golang/projects/vigilate/config"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

	"github.com/pusher/pusher-http-go"
)

type PusherClient struct {
	Client    pusher.Client
	Prefs     *helpers.Preferences
	Scheduler *Scheduler
}

func NewPusherClient(cfg *config.Config, Prefs *helpers.Preferences, scheduler *Scheduler) *PusherClient {

	Prefs.SetPref("pusher-host", cfg.PusherConfig.Host)
	Prefs.SetPref("pusher-port", cfg.PusherConfig.Port)
	Prefs.SetPref("pusher-key", cfg.PusherConfig.Key)

	// create pusher client
	wsClient := pusher.Client{
		AppID:  cfg.PusherConfig.App,
		Secret: cfg.PusherConfig.Secret,
		Key:    cfg.PusherConfig.Key,
		Secure: cfg.PusherConfig.Secure,
		Host:   fmt.Sprintf("%s:%s", cfg.PusherConfig.Host, cfg.PusherConfig.Port),
	}

	return &PusherClient{
		Client:    wsClient,
		Prefs:     Prefs,
		Scheduler: scheduler,
	}
}

func (p *PusherClient) BroadcastMessage(channel string, messageType string, data map[string]string) {
	err := p.Client.Trigger(channel, messageType, data)
	if err != nil {
		log.Println(err)
	}
}

func (p *PusherClient) PushStatusChangedEvent(h models.Host, hs models.HostService, newStatus string) {
	data := map[string]string{
		"message":         fmt.Sprintf("%s on %s reports %s", hs.Service.ServiceName, hs.HostName, newStatus),
		"host_id":         strconv.Itoa(hs.HostID),
		"host_service_id": strconv.Itoa(hs.ID),
		"host_name":       hs.HostName,
		"service_name":    hs.Service.ServiceName,
		"icon":            hs.Service.Icon,
		"status":          newStatus,
		"last_check":      time.Now().Format("2006-01-02 3:04:06 PM"),
	}
	p.BroadcastMessage("public-channel", "host-service-status-changed", data)
}

func (p *PusherClient) PushScheduleChangedEvent(hs models.HostService, newStatus string) {

	yearOne := time.Date(0001, 1, 1, 0, 0, 0, 1, time.UTC)
	data := map[string]string{
		"host_service_id": strconv.Itoa(hs.ID),
		"service_id":      strconv.Itoa(hs.ServiceID),
		"host_id":         strconv.Itoa(hs.HostID),
		"last_run":        time.Now().Format("2006-01-02 3:04:05 PM"),
		"host":            hs.HostName,
		"service":         hs.Service.ServiceName,
		"schedule":        fmt.Sprintf("@every %d%s", hs.ScheduleNumber, hs.ScheduleUnit),
		"status":          newStatus,
		"icon":            hs.Service.Icon,
	}

	if p.Scheduler.Cron.Entry(p.Scheduler.MonitorMap[hs.ID]).Next.After(yearOne) {
		data["next_run"] = p.Scheduler.Cron.Entry(p.Scheduler.MonitorMap[hs.ID]).Next.Format("2006-01-02 3:04:05 PM")
	} else {
		data["next_run"] = "Pending..."
	}

	p.BroadcastMessage("public-channel", "schedule-changed-event", data)
}
