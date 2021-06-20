package services

import (
	"log"
	"strconv"
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"
)

type Job struct {
	HostServiceID int
	Repo          repository.DatabaseRepo
	Prefs         *helpers.Preferences
	WsClient      *PusherClient
	Tester        *Tester
}

func (j Job) Run() {
	log.Println("******** Running check for", j.HostServiceID)
	hs, err := j.Repo.GetHostServiceByID(j.HostServiceID)
	if err != nil {
		log.Println(err)
		return
	}

	h, err := j.Repo.GetHostByID(hs.HostID)
	if err != nil {
		log.Println(err)
		return
	}

	msg, newStatus := j.Tester.Test(h, hs)
	if newStatus != hs.Status {
		j.updateHostServiceStatusCount(h, hs, newStatus, msg)
	}
}

func (j Job) updateHostServiceStatusCount(h models.Host, hs models.HostService, newStatus string, msg string) {
	// update host service in the db with status (if changed) and last check
	hs.Status = newStatus
	hs.LastCheck = time.Now()
	hs.LastMessage = msg
	err := j.Repo.UpdateHostService(hs)
	if err != nil {
		log.Println(err)
		return
	}

	pending, healthy, warning, problem, err := j.Repo.GetAllServiceStatusCounts()
	if err != nil {
		log.Println(err)
		return
	}

	data := map[string]string{
		"pending_count": strconv.Itoa(pending),
		"healthy_count": strconv.Itoa(healthy),
		"warning_count": strconv.Itoa(warning),
		"problem_count": strconv.Itoa(problem),
	}

	j.WsClient.BroadcastMessage("public-channel", "host-service-count-changed", data)

	// if appropriate, send email or sms message
	log.Println("new status:", newStatus, "msg:", msg)
}
