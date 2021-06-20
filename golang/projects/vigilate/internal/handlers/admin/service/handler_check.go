package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
	"github.com/go-chi/chi/v5"
)

type HandlerTestCheckNow struct {
	repo     Repository
	tester   *services.Tester
	wsClient *services.PusherClient
}

func NewHandlerTestCheck(repo Repository, tester *services.Tester, wsClient *services.PusherClient) *HandlerTestCheckNow {
	return &HandlerTestCheckNow{repo: repo, tester: tester, wsClient: wsClient}
}

func (h *HandlerTestCheckNow) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostServiceID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	oldStatus := chi.URLParam(r, "oldStatus")
	ok := true
	log.Println("hostServiceID:", hostServiceID, "oldStatus:", oldStatus)

	hs, err := h.repo.GetHostServiceByID(hostServiceID)
	if err != nil {
		log.Println(err)
		ok = false
	}

	host, err := h.repo.GetHostByID(hs.HostID)
	if err != nil {
		log.Println(err)
		ok = false
	}

	message, newStatus := h.tester.Test(host, hs)
	log.Println("message:", message, "newStatus:", newStatus)

	// if the host service status has changed, broadcast to all clients
	if newStatus != hs.Status {
		h.wsClient.PushStatusChangedEvent(host, hs, newStatus)
		err := h.repo.InsertEvent(models.Event{
			EventType:     newStatus,
			HostServiceID: hs.ID,
			HostID:        host.ID,
			ServiceName:   hs.Service.ServiceName,
			HostName:      hs.HostName,
			Message:       message,
		})
		if err != nil {
			log.Println(err)
		}
	}

	// broadcast service status changed event
	if newStatus != hs.Status {
		h.wsClient.PushStatusChangedEvent(host, hs, newStatus)
	}
	// update the host service in the database with status (if changed) and last check
	hs.Status = newStatus
	hs.LastCheck = time.Now()
	hs.LastMessage = message
	err = h.repo.UpdateHostService(hs)
	if err != nil {
		log.Println(err)
		return
	}

	var resp models.TestCheckResp
	if ok {
		resp = models.TestCheckResp{
			OK:            true,
			Message:       message,
			ServiceID:     hs.ServiceID,
			HostServiceID: hs.ID,
			HostID:        hs.HostID,
			OldStatus:     oldStatus,
			NewStatus:     newStatus,
			LastCheck:     time.Now(),
		}
	} else {
		resp.OK = false
		resp.Message = "Something went wrong"
	}

	marshal, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}
