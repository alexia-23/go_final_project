package utils

import (
	"encoding/json"
	"github.com/alexia-23/go_final_project/pkg/domain"
	"github.com/alexia-23/go_final_project/pkg/logger"
	"net/http"
)

func WriteError(w http.ResponseWriter, msg string, code int) {
	log := logger.Get()
	log.Printf("ERROR: %s", msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := domain.TaskCreateResponse{Error: &msg}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
