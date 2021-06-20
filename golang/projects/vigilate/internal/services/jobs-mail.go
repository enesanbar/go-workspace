package services

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/aymerick/douceur/inliner"
	mail "github.com/xhit/go-simple-mail/v2"
	"jaytaylor.com/html2text"
)

// NewWorker takes a numeric id and a channel w/ worker pool.
func NewWorker(id int, workerPool chan chan models.MailJob, prefs *helpers.Preferences) Worker {
	return Worker{
		id:         id,
		jobQueue:   make(chan models.MailJob),
		workerPool: workerPool,
		quitChan:   make(chan bool),
		prefs:      prefs,
	}
}

// Worker holds info for a pool worker
type Worker struct {
	id         int
	jobQueue   chan models.MailJob
	workerPool chan chan models.MailJob
	quitChan   chan bool
	prefs      *helpers.Preferences
}

// start starts the worker
func (w Worker) start() {
	go func() {
		for {
			// Add jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				w.processMailQueueJob(job.MailMessage)
			case <-w.quitChan:
				fmt.Printf("worker%d stopping\n", w.id)
				return
			}
		}
	}()
}

// stop the worker
func (w Worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}

type DispatcherConfig struct {
	jobQueue   chan models.MailJob
	maxWorkers int
	Prefs      *helpers.Preferences
}

func NewDispatcherConfig(preferences *helpers.Preferences) *DispatcherConfig {
	mailQueue := make(chan models.MailJob, 5)
	return &DispatcherConfig{jobQueue: mailQueue, maxWorkers: 5, Prefs: preferences}
}

// NewDispatcher creates, and returns a new Dispatcher object.
func NewDispatcher(cfg *DispatcherConfig) helpers.Runnable {
	workerPool := make(chan chan models.MailJob, cfg.maxWorkers)
	return &Dispatcher{
		jobQueue:   cfg.jobQueue,
		maxWorkers: cfg.maxWorkers,
		workerPool: workerPool,
	}
}

// Dispatcher holds info for a dispatcher
type Dispatcher struct {
	workerPool chan chan models.MailJob
	maxWorkers int
	jobQueue   chan models.MailJob
	prefs      *helpers.Preferences
}

func (d *Dispatcher) Stop() error {
	return nil
}

// Start runs the workers
func (d *Dispatcher) Start() error {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(i+1, d.workerPool, d.prefs)
		worker.start()
	}

	go d.dispatch()
	return nil
}

// dispatch dispatches worker
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func() {
				workerJobQueue := <-d.workerPool
				workerJobQueue <- job
			}()
		}
	}
}

// processMailQueueJob processes the main queue Job (sends email)
func (w Worker) processMailQueueJob(mailMessage models.MailData) {
	log.Println("processing mail")
	data := struct {
		Content       template.HTML
		From          string
		FromName      string
		PreferenceMap map[string]string
		IntMap        map[string]int
		StringMap     map[string]string
		FloatMap      map[string]float32
		RowSets       map[string]interface{}
	}{
		Content:   mailMessage.Content,
		FromName:  mailMessage.FromName,
		From:      mailMessage.FromAddress,
		IntMap:    mailMessage.IntMap,
		StringMap: mailMessage.StringMap,
		FloatMap:  mailMessage.FloatMap,
		RowSets:   mailMessage.RowSets,
	}

	paths := []string{
		"./views/mail.tmpl",
	}

	t := template.Must(template.New("mail.tmpl").ParseFiles(paths...))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		fmt.Print(err)
	}

	result := tpl.String()

	plainText, err := html2text.FromString(result, html2text.Options{PrettyTables: true})
	if err != nil {
		plainText = ""
	}

	var formattedMessage string

	formattedMessage, err = inliner.Inline(result)
	if err != nil {
		log.Println(err)
		formattedMessage = result
	}

	port, _ := strconv.Atoi(w.prefs.GetPref("smtp_port"))

	server := mail.NewSMTPClient()
	server.Host = w.prefs.GetPref("smtp_server")
	server.Port = port
	server.Username = w.prefs.GetPref("smtp_user")
	server.Password = w.prefs.GetPref("smtp_password")
	if w.prefs.GetPref("smtp_server") == "localhost" {
		server.Authentication = mail.AuthPlain
	} else {
		server.Authentication = mail.AuthLogin
	}
	server.Encryption = mail.EncryptionTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		log.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(mailMessage.FromAddress).
		AddTo(mailMessage.ToAddress).
		SetSubject(mailMessage.Subject)

	if len(mailMessage.AdditionalTo) > 0 {
		for _, x := range mailMessage.AdditionalTo {
			email.AddTo(x)
		}
	}

	if len(mailMessage.CC) > 0 {
		for _, x := range mailMessage.CC {
			email.AddCc(x)
		}
	}

	if len(mailMessage.Attachments) > 0 {
		for _, x := range mailMessage.Attachments {
			email.AddAttachment(x)
		}
	}

	email.SetBody(mail.TextHTML, formattedMessage)
	email.AddAlternative(mail.TextPlain, plainText)

	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}
}
