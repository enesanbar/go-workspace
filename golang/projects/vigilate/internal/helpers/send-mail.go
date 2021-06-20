package helpers

import (
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
)

// SendEmail sends an email
func SendEmail(mailMessage models.MailData, prefs *Preferences) {
	if mailMessage.FromAddress == "" {
		mailMessage.FromAddress = prefs.GetPref("smtp_from_email")
		mailMessage.FromName = prefs.GetPref("smtp_from_name")
	}

	//job := channeldata.MailJob{MailMessage: mailMessage}
	//app.MailQueue <- job
}
