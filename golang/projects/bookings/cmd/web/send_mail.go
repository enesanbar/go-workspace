package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/enesanbar/workspace/projects/bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func ListenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMessage(msg)
		}
	}()
}

func sendMessage(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)

	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		data, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
		if err != nil {
			app.ErrorLog.Println(err)
		}

		mailTemplate := string(data)
		renderedTemplate := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, renderedTemplate)
	}

	err = email.Send(client)
	if err != nil {
		app.ErrorLog.Println(err)
	} else {
		app.InfoLog.Println("email send")
	}

}
