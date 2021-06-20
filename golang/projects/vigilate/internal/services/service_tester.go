package services

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services/certificateutils"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
)

const (
	HTTP           = 1
	HTTPS          = 2
	SSLCertificate = 3
)

type Tester struct {
	WsClient *PusherClient
	Repo     repository.DatabaseRepo
	Prefs    *helpers.Preferences
}

func NewTester(wsClient *PusherClient, repo repository.DatabaseRepo, prefs *helpers.Preferences) *Tester {
	return &Tester{WsClient: wsClient, Repo: repo, Prefs: prefs}
}

func (t *Tester) Test(h models.Host, hs models.HostService) (string, string) {
	var msg, newStatus string

	switch hs.ServiceID {
	case HTTP:
		msg, newStatus = testHTTPForHost(h.URL)
	case HTTPS:
		msg, newStatus = testHTTPSForHost(h.URL)
	case SSLCertificate:
		msg, newStatus = testSSLForHost(h.URL)
	}

	// if the host service status has changed, broadcast to all clients
	if newStatus != hs.Status {
		t.WsClient.PushStatusChangedEvent(h, hs, newStatus)
		err := t.Repo.InsertEvent(models.Event{
			EventType:     newStatus,
			HostServiceID: hs.ID,
			HostID:        h.ID,
			ServiceName:   hs.Service.ServiceName,
			HostName:      hs.HostName,
			Message:       msg,
		})
		if err != nil {
			log.Println(err)
		}

		// send email if appropriate
		if t.Prefs.GetPref("notify_via_email") == "1" && hs.Status != "pending" {
			data := models.MailData{
				ToName:    t.Prefs.GetPref("notify_name"),
				ToAddress: t.Prefs.GetPref("notify_email"),
			}

			if newStatus == "healthy" {
				data.Subject = fmt.Sprintf("HEALTHY: service %s on %s", hs.Service.ServiceName, hs.HostName)
				data.Content = template.HTML(fmt.Sprintf(`
					<p>Service %s on %s reported healthy status</p>
					<p><strong>Message received %s</strong</p>
				`, hs.Service.ServiceName, hs.HostName, msg))
			} else if newStatus == "problem" {
				data.Subject = fmt.Sprintf("PROBLEM: service %s on %s", hs.Service.ServiceName, hs.HostName)
				data.Content = template.HTML(fmt.Sprintf(`
					<p>Service %s on %s reported problem status</p>
					<p><strong>Message received %s</strong</p>
				`, hs.Service.ServiceName, hs.HostName, msg))
			} else if newStatus == "warning" {

			}

			helpers.SendEmail(data, t.Prefs)
		}
		// send sms if appropriate
		if t.Prefs.GetPref("notify_via_sms") == "1" {
			to := t.Prefs.GetPref("sms_notify_number")
			body := ""
			if newStatus == "healthy" {
				body = fmt.Sprintf("Service %s on %s is healthy", hs.Service.ServiceName, hs.HostName)
			} else if newStatus == "problem" {
				body = fmt.Sprintf("Service %s on %s reports a problem %s", hs.Service.ServiceName, hs.HostName, msg)
			} else if newStatus == "warning" {
				body = fmt.Sprintf("Service %s on %s reports a warning %s", hs.Service.ServiceName, hs.HostName, msg)
			}

			err := helpers.SendTextTwilio(to, body, t.Prefs)
			if err != nil {
				log.Println(err)
			}
		}
	}

	t.WsClient.PushScheduleChangedEvent(hs, newStatus)
	return msg, newStatus

}

func testHTTPForHost(url string) (string, string) {
	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}

	url = strings.Replace(url, "https://", "http://", -1)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("%s - %s", url, "error connecting"), "problem"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("%s - %s", url, resp.Status), "problem"
	}

	return fmt.Sprintf("%s - %s", url, resp.Status), "healthy"
}

func testHTTPSForHost(url string) (string, string) {
	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}

	url = strings.Replace(url, "http://", "https://", -1)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("%s - %s", url, "error connecting"), "problem"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("%s - %s", url, resp.Status), "problem"
	}

	return fmt.Sprintf("%s - %s", url, resp.Status), "healthy"
}

// scanHost gets cert details from an internet host
func scanHost(hostname string, certDetailsChannel chan certificateutils.CertificateDetails, errorsChannel chan error) {

	res, err := certificateutils.GetCertificateDetails(hostname, 10)
	if err != nil {
		errorsChannel <- err
	} else {
		certDetailsChannel <- res
	}
}

func testSSLForHost(url string) (string, string) {
	if strings.HasPrefix(url, "https://") {
		url = strings.Replace(url, "https://", "", -1)
	}

	if strings.HasPrefix(url, "http://") {
		url = strings.Replace(url, "http://", "", -1)
	}

	// scanning ssl cert for expiry date
	var certDetailsChannel chan certificateutils.CertificateDetails
	var errorsChannel chan error
	certDetailsChannel = make(chan certificateutils.CertificateDetails, 1)
	errorsChannel = make(chan error, 1)

	var msg, newStatus string
	scanHost(url, certDetailsChannel, errorsChannel)

	for i, certDetailsInQueue := 0, len(certDetailsChannel); i < certDetailsInQueue; i++ {
		certDetails := <-certDetailsChannel
		certificateutils.CheckExpirationStatus(&certDetails, 30)

		if certDetails.ExpiringSoon {

			if certDetails.DaysUntilExpiration < 7 {
				msg = certDetails.Hostname + " expiring in " + strconv.Itoa(certDetails.DaysUntilExpiration) + " days"
				newStatus = "problem"
			} else {
				msg = certDetails.Hostname + " expiring in " + strconv.Itoa(certDetails.DaysUntilExpiration) + " days"
				newStatus = "warning"
			}

		} else {
			msg = certDetails.Hostname + " expiring in " + strconv.Itoa(certDetails.DaysUntilExpiration) + " days"
			newStatus = "healthy"
		}

	}

	return msg, newStatus
}
