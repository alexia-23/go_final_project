package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexia-23/go_final_project/pkg/domain"
)

func (s *Server) PutTask(w http.ResponseWriter, r *http.Request) {

	var payload domain.Task
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		msg := "invalid JSON: " + err.Error()
		s.WriteError(w, msg, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if payload.ID == 0 {
		s.WriteError(w, "missing field: id", http.StatusBadRequest)
		return
	}

	if payload.Title == "" {
		s.WriteError(w, "missing field: title", http.StatusBadRequest)
		return
	}

	if !ValidateRepeat(payload.Repeat) {
		s.WriteError(w, "invalid field: repeat", http.StatusBadRequest)
		return
	}

	now := time.Now()
	today := now.Format(DATE_FORMAT)
	if payload.Date == "" {
		payload.Date = today
	} else {
		_, err := time.Parse(DATE_FORMAT, payload.Date)
		if err != nil {
			s.WriteError(w, "invalid field: date", http.StatusBadRequest)
			return
		}
	}

	if payload.Repeat == "" && today > payload.Date {
		payload.Date = today
	} else if today > payload.Date {
		next, err := NextDate(now, payload.Date, payload.Repeat)
		if err != nil {
			s.WriteError(w, "invalid field: date", http.StatusBadRequest)
			return
		}
		payload.Date = next
	}

	id, err := s.Storage.UpdateTask(payload)
	if err != nil {
		s.WriteError(w, "error saving into db: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := domain.TaskCreateResponse{ID: &id}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
