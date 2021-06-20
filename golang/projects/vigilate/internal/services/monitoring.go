package services

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"
)

type Monitoring struct {
	prefs      *helpers.Preferences
	wsClient   *PusherClient
	repository repository.DatabaseRepo
	Scheduler  *Scheduler
	tester     *Tester
}

func NewRunnableMonitoring(prefs *helpers.Preferences, wsClient *PusherClient, repository repository.DatabaseRepo, scheduler *Scheduler, tester *Tester) helpers.Runnable {
	return &Monitoring{prefs: prefs, wsClient: wsClient, repository: repository, Scheduler: scheduler, tester: tester}
}

func NewMonitoring(prefs *helpers.Preferences, wsClient *PusherClient, repository repository.DatabaseRepo, scheduler *Scheduler, tester *Tester) *Monitoring {
	return &Monitoring{prefs: prefs, wsClient: wsClient, repository: repository, Scheduler: scheduler, tester: tester}
}

func (m *Monitoring) Start() error {
	if m.prefs.GetPref("monitoring_live") == "0" {
		return nil
	}

	m.Scheduler.Start()

	// trigger a message to broadcast to all clients that app is starting to monitor
	data := make(map[string]string)
	data["message"] = "Monitoring is starting..."
	err := m.wsClient.Client.Trigger("public-channel", "app-starting", data)
	if err != nil {
		log.Println(err)
	}

	// get all the services that we want to monitor
	servicesToMonitor, err := m.repository.GetServicesToMonitor()
	if err != nil {
		log.Println(err)
	}

	// range through the services
	for _, service := range servicesToMonitor {
		log.Println("service to monitor:", service.HostName, service.Service.ServiceName)

		// get the schedule unit and number
		var schedule string
		if service.ScheduleUnit == "d" {
			schedule = fmt.Sprintf("@every %d%s", service.ScheduleNumber*24, "h")
		} else {
			schedule = fmt.Sprintf("@every %d%s", service.ScheduleNumber, service.ScheduleUnit)
		}

		// create a job
		job := Job{
			HostServiceID: service.ID,
			Repo:          m.repository,
			Prefs:         m.prefs,
			WsClient:      m.wsClient,
			Tester:        m.tester,
		}

		scheduleID, err := m.Scheduler.Cron.AddJob(schedule, job)
		if err != nil {
			log.Println(err)
		}

		// save the id of the job, so we can start/stop it
		m.Scheduler.MonitorMap[service.ID] = scheduleID

		yearOne := time.Date(0001, 11, 17, 20, 34, 58, 65138737, time.UTC)

		// broadcast over websockets the fact that the service is scheduled
		payload := map[string]string{
			"message":         "scheduling",
			"host_service_id": strconv.Itoa(service.ID),
			"host":            service.HostName,
			"service":         service.Service.ServiceName,
			"schedule":        fmt.Sprintf("@every %d%s", service.ScheduleNumber, service.ScheduleUnit),
		}

		if m.Scheduler.Cron.Entry(m.Scheduler.MonitorMap[service.ID]).Next.After(yearOne) {
			payload["next_run"] = m.Scheduler.Cron.Entry(m.Scheduler.MonitorMap[service.ID]).Next.Format("2006-01-02 3:04:05 PM")
		} else {
			payload["next_run"] = "Pending..."
		}

		if service.LastCheck.After(yearOne) {
			payload["last_run"] = service.LastCheck.Format("2006-01-02 3:04:05 PM")
		} else {
			payload["last_run"] = "Pending..."
		}

		err = m.wsClient.Client.Trigger("public-channel", "next-run-event", payload)
		if err != nil {
			log.Println(err)
		}

		err = m.wsClient.Client.Trigger("public-channel", "schedule-changed-event", payload)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (m *Monitoring) Stop() error {
	return m.Scheduler.Stop()
}
