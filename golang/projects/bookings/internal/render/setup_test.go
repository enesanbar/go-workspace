package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/enesanbar/workspace/projects/bookings/internal/config"
	"github.com/enesanbar/workspace/projects/bookings/internal/models"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false
	testApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session
	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct {
}

func (m myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (m myWriter) Write(bytes []byte) (int, error) {
	l := len(bytes)
	return l, nil
}

func (m myWriter) WriteHeader(statusCode int) {
}
