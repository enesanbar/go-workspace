package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/config"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/repository"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

type Session struct {
	Manager *scs.SessionManager
}

func NewSessionStore(conn *repository.DBConnection) scs.Store {
	return postgresstore.New(conn.Conn)
}

func NewSession(session *scs.SessionManager, sessionStore scs.Store, cfg *config.Config) *Session {
	session.Store = sessionStore
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Name = fmt.Sprintf("gbsession_id_%s", cfg.Identifier)
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = cfg.InProduction

	return &Session{Manager: session}
}

// IsAuthenticated returns true if a user is authenticated
func (s *Session) IsAuthenticated(r *http.Request) bool {
	exists := s.Manager.Exists(r.Context(), "userID")
	return exists
}
