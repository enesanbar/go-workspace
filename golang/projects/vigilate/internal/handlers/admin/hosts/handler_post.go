package hosts

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"
	"github.com/go-chi/chi/v5"
)

type HandlerPost struct {
	Repository Repository
	Renderer   *services.Renderer
	Session    *session.Session
}

func NewHandlerPost(repository Repository, renderer *services.Renderer, session *session.Session) *HandlerPost {
	return &HandlerPost{Repository: repository, Renderer: renderer, Session: session}
}

// ServeHTTP handles posting of host form
func (h *HandlerPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var host models.Host

	if id > 0 {
		// get existing host
		h, err := h.Repository.GetHostByID(id)
		if err != nil {
			log.Println(err)
			return
		}
		host = h
	}

	host.HostName = r.Form.Get("host_name")
	host.CanonicalName = r.Form.Get("canonical_name")
	host.URL = r.Form.Get("url")
	host.IP = r.Form.Get("ip")
	host.IPV6 = r.Form.Get("ipv6")
	host.Location = r.Form.Get("location")
	host.OS = r.Form.Get("os")

	active, _ := strconv.Atoi(r.Form.Get("active"))
	host.Active = active

	if id > 0 {
		err := h.Repository.UpdateHost(host)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		newID, err := h.Repository.InsertHost(host)
		if err != nil {
			log.Println(err)
			h.Renderer.ServerError(w, r, err)
			return
		}
		host.ID = newID
	}
	h.Session.Manager.Put(r.Context(), "flash", "changes saved")
	http.Redirect(w, r, fmt.Sprintf("/admin/host/%d", host.ID), http.StatusSeeOther)
}
