package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/enesanbar/workspace/golang/projects/vigilate/config"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"
)

type HandlerLogout struct {
	repo    Repository
	cfg     *config.Config
	session *session.Session
}

func NewHandlerLogout(repo Repository, cfg *config.Config, session *session.Session) *HandlerLogout {
	return &HandlerLogout{repo: repo, cfg: cfg, session: session}
}

// ServeHTTP logs the user out
func (h *HandlerLogout) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// delete the remember me token, if any
	cookie, err := r.Cookie(fmt.Sprintf("_%s_gowatcher_remember", h.cfg.Identifier))
	if err != nil {
	} else {
		key := cookie.Value
		// have a remember token, so get the token
		if len(key) > 0 {
			// key length > 0, so it might be a valid token
			split := strings.Split(key, "|")
			hash := split[1]
			err = h.repo.DeleteToken(hash)
			if err != nil {
				log.Println(err)
			}
		}
	}

	// delete the remember me cookie, if any
	delCookie := http.Cookie{
		Name:     fmt.Sprintf("_%s_gowatcher_remember", h.cfg.Identifier),
		Value:    "",
		Domain:   h.cfg.Domain,
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}
	http.SetCookie(w, &delCookie)

	_ = h.session.Manager.RenewToken(r.Context())
	_ = h.session.Manager.Destroy(r.Context())
	_ = h.session.Manager.RenewToken(r.Context())

	h.session.Manager.Put(r.Context(), "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
