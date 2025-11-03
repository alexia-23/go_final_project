package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/alexia-23/go_final_project/pkg/db"
	"github.com/alexia-23/go_final_project/pkg/logger"
	"github.com/alexia-23/go_final_project/pkg/utils"
	"net/http"
	"strconv"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()
	if r.Method != http.MethodGet {
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
	log.Println("task selected: ", task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
