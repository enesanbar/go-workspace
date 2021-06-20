package main

import (
	"encoding/json"
	"net/http"
)

func (a *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	marshal, err := json.Marshal(wrapper)
	if err != nil {
		a.logger.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(marshal)
	if err != nil {
		a.logger.Println(err)
	}
}

func (a *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	type jsonError struct {
		Message string `json:"message"`
	}

	resp := jsonError{
		Message: err.Error(),
	}

	a.writeJSON(w, statusCode, resp, "error")

}
