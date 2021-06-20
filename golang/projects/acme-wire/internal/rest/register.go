package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/modules/data"
)

// RegisterModel will validate and save a registration
//go:generate mockery -name=RegisterModel
type RegisterModel interface {
	Do(ctx context.Context, in *data.Person) (int, error)
}

// RegisterHandler is the HTTP handler for the "Register" endpoint
type RegisterHandler struct {
	registerer RegisterModel
}

func NewRegisterHandler(registerer RegisterModel) *RegisterHandler {
	return &RegisterHandler{registerer: registerer}
}

// ServeHTTP implements http.Handler
func (h *RegisterHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// set latency budget for this API
	subCtx, cancel := context.WithTimeout(request.Context(), 1500*time.Millisecond)
	defer cancel()

	// extract payload from request
	requestPayload, err := h.extractPayload(request)
	if err != nil {
		// output error
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// register person
	id, err := h.register(subCtx, requestPayload)
	if err != nil {
		// not need to log here as we can expect other layers to do so
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// happy path
	response.Header().Add("Location", fmt.Sprintf("/person/%d/", id))
	response.WriteHeader(http.StatusCreated)
}

// extract payload from request
func (h *RegisterHandler) extractPayload(request *http.Request) (*registerRequest, error) {
	requestPayload := &registerRequest{}
	err := json.NewDecoder(request.Body).Decode(requestPayload)
	return requestPayload, err
}

// call the logic layer
func (h *RegisterHandler) register(ctx context.Context, requestPayload *registerRequest) (int, error) {
	person := &data.Person{
		FullName: requestPayload.FullName,
		Phone:    requestPayload.Phone,
		Currency: requestPayload.Currency,
	}

	return h.registerer.Do(ctx, person)
}

// register endpoint request format
type registerRequest struct {
	// FullName of the person
	FullName string `json:"fullName"`
	// Phone of the person
	Phone string `json:"phone"`
	// Currency the wish to register in
	Currency string `json:"currency"`
}
