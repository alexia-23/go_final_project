package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alexia-23/go_final_project/pkg/domain"
)

func (s *Server) WriteError(w http.ResponseWriter, msg string, code int) {
	log.Printf("ERROR: %s", msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := domain.TaskCreateResponse{Error: &msg}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
