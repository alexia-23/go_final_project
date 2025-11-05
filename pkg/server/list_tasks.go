package server

import (
	"encoding/json"
	"net/http"

	"github.com/alexia-23/go_final_project/pkg/domain"
)

func (s *Server) ListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.WriteError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	tasks, err := s.Storage.SelectTasks()
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := domain.TaskListResponse{Tasks: tasks}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
