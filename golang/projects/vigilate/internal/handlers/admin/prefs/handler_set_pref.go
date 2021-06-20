package prefs

import (
	"encoding/json"
	"net/http"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
)

type HandlerSetPref struct {
	repo  Repository
	prefs *helpers.Preferences
}

func NewHandlerSetPref(repo Repository, prefs *helpers.Preferences) *HandlerSetPref {
	return &HandlerSetPref{repo: repo, prefs: prefs}
}

func (h *HandlerSetPref) SetSystemPref(w http.ResponseWriter, r *http.Request) {
	prefName := r.PostForm.Get("pref_name")
	prefValue := r.PostForm.Get("pref_value")

	resp := models.TestCheckResp{
		OK:      true,
		Message: "",
	}

	err := h.repo.UpdateSystemPref(prefName, prefValue)
	if err != nil {
		resp.OK = false
		resp.Message = err.Error()
	}

	h.prefs.SetPref(prefName, prefValue)

	marshal, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}
