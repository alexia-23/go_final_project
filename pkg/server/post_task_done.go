package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) PostTaskDone(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.WriteError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	params := r.URL.Query()
	raw := params.Get("id")
	id, err := strconv.Atoi(raw)
	if err != nil {
		s.WriteError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	task, err := s.Storage.SelectTaskById(id)
	if err == sql.ErrNoRows {
		s.WriteError(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if task.Repeat == "" {
		_, err = s.Storage.DeleteTaskById(id)
	} else {
		task.Date, err = NextDate(time.Now(), task.Date, task.Repeat)
		_, err = s.Storage.UpdateTask(task)
	}

	if err == sql.ErrNoRows {
		s.WriteError(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]any{}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
