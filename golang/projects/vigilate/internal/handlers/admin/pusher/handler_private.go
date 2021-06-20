package pusher

import (
	"fmt"
	"log"
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
)

type HandlerPrivate struct {
	wsClient *services.PusherClient
}

func NewHandlerPrivate(wsClient *services.PusherClient) *HandlerPrivate {
	return &HandlerPrivate{wsClient: wsClient}
}

func (h *HandlerPrivate) SendPrivateMessage(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	id := r.URL.Query().Get("id")

	data := map[string]string{
		"message": msg,
	}

	err := h.wsClient.Client.Trigger(fmt.Sprintf("private-channel-%s", id), "private-message", data)
	if err != nil {
		log.Println(err)
		return
	}
}
