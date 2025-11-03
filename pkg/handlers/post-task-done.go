package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/alexia-23/go_final_project/pkg/db"
	"github.com/alexia-23/go_final_project/pkg/logger"
	"github.com/alexia-23/go_final_project/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

func PostTaskDone(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	params := r.URL.Query()
	raw := params.Get("id")
	id, err := strconv.Atoi(raw)
	if err != nil {
		utils.WriteError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	task, err := db.SelectTaskById(id)
	if err == sql.ErrNoRows {
		utils.WriteError(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	if err != nil {
		utils.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if task.Repeat == "" {
		_, err = db.DeleteTaskById(id)
	} else {
		task.Date, err = utils.NextDate(time.Now(), task.Date, task.Repeat)
		_, err = db.UpdateTask(task)
	}

	if err == sql.ErrNoRows {
		utils.WriteError(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	if err != nil {
		utils.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("task marked done: ", task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]any{}); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
