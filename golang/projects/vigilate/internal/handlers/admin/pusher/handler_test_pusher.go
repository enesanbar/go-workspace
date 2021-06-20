package pusher

import (
	"log"
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
)

type HandlerTestPusher struct {
	wsClient *services.PusherClient
}

func NewHandlerTestPusher(wsClient *services.PusherClient) *HandlerTestPusher {
	return &HandlerTestPusher{wsClient: wsClient}
}

func (h *HandlerTestPusher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["message"] = "hello"

	err := h.wsClient.Client.Trigger("public-channel", "test-event", data)
	if err != nil {
		log.Println(err)
	}
}
