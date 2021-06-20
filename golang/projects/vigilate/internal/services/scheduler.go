package services

import (
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	Cron       *cron.Cron
	MonitorMap map[int]cron.EntryID
	Prefs      *helpers.Preferences
}

func NewScheduler(prefs *helpers.Preferences) *Scheduler {
	localZone, _ := time.LoadLocation("Local")
	c := cron.New(
		cron.WithLocation(localZone),
		cron.WithChain(
			cron.DelayIfStillRunning(cron.DefaultLogger),
			cron.Recover(cron.DefaultLogger),
		),
	)
	s := &Scheduler{
		Cron:       c,
		MonitorMap: make(map[int]cron.EntryID),
		Prefs:      prefs,
	}
	return s
}

func (s Scheduler) Start() error {
	if s.Prefs.GetPref("monitoring_live") == "1" {
		s.Cron.Start()
	}
	return nil
}

func (s Scheduler) Stop() error {
	s.Cron.Stop()
	return nil
}
