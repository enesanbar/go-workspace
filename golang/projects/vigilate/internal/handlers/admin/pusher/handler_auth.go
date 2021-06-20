package pusher

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"
	"github.com/pusher/pusher-http-go"
)

type HandlerAuth struct {
	session  *session.Session
	repo     Repository
	wsClient *services.PusherClient
}

func NewHandlerAuth(session *session.Session, repo Repository, wsClient *services.PusherClient) *HandlerAuth {
	return &HandlerAuth{session: session, repo: repo, wsClient: wsClient}
}

func (h *HandlerAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID := h.session.Manager.GetInt(r.Context(), "userID")
	user, err := h.repo.GetUserById(userID)
	if err != nil {
		log.Println(err)
	}

	params, _ := io.ReadAll(r.Body)
	presenceData := pusher.MemberData{
		UserID: strconv.Itoa(userID),
		UserInfo: map[string]string{
			"name": user.FirstName,
			"id":   strconv.Itoa(userID),
		},
	}

	resp, err := h.wsClient.Client.AuthenticatePresenceChannel(params, presenceData)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resp)
}
