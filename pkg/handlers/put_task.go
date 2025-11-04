package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexia-23/go_final_project/pkg/db"
	"github.com/alexia-23/go_final_project/pkg/domain"
	"github.com/alexia-23/go_final_project/pkg/logger"
	"github.com/alexia-23/go_final_project/pkg/utils"
)

func PutTask(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()
	if r.Method != http.MethodPut {
		utils.WriteError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var payload domain.Task
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		msg := "invalid JSON: " + err.Error()
		utils.WriteError(w, msg, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if payload.Id == 0 {
		utils.WriteError(w, "missing field: id", http.StatusBadRequest)
		return
	}

	if payload.Title == "" {
		utils.WriteError(w, "missing field: title", http.StatusBadRequest)
		return
	}

	if !utils.ValidateRepeat(payload.Repeat) {
		utils.WriteError(w, "invalid field: repeat", http.StatusBadRequest)
		return
	}

	now := time.Now()
	today := now.Format("20060102")
	if payload.Date == "" {
		payload.Date = today
	} else {
		_, err := time.Parse("20060102", payload.Date)
		if err != nil {
			utils.WriteError(w, "invalid field: date", http.StatusBadRequest)
			return
		}
	}

	if payload.Repeat == "" && today > payload.Date {
		payload.Date = today
	} else if today > payload.Date {
		next, err := utils.NextDate(now, payload.Date, payload.Repeat)
		if err != nil {
			utils.WriteError(w, "invalid field: date", http.StatusBadRequest)
			return
		}
		payload.Date = next
	}

	id, err := db.UpdateTask(payload)
	if err != nil {
		utils.WriteError(w, "error saving into db: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("task updated: ", id)

	resp := domain.TaskCreateResponse{ID: &id}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
