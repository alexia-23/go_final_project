package handlers

import (
	"encoding/json"
	"github.com/alexia-23/go_final_project/pkg/db"
	"github.com/alexia-23/go_final_project/pkg/domain"
	"github.com/alexia-23/go_final_project/pkg/logger"
	"github.com/alexia-23/go_final_project/pkg/utils"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload domain.TaskCreatePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		msg := "invalid JSON: " + err.Error()
		utils.WriteError(w, msg, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if payload.Title == "" {
		utils.WriteError(w, "missing field: title", http.StatusBadRequest)
		return
	}

	id, err := db.InsertTask(payload)
	if err != nil {
		utils.WriteError(w, "error saving into db: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("task created: ", id)

	resp := domain.TaskCreateResponse{ID: &id}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
