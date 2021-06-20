package main

import (
	"encoding/json"
	"net/http"
)

func (a *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	status := AppStatus{
		Status:      "Available",
		Environment: a.config.env,
		Version:     version,
	}

	marshal, err := json.Marshal(status)
	if err != nil {
		a.logger.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(marshal)
	if err != nil {
		a.logger.Println(err)
	}
}
