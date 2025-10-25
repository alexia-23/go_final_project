package handlers

import (
	"encoding/json"
	"github.com/alexia-23/go_final_project/pkg/db"
	"github.com/alexia-23/go_final_project/pkg/domain"
	"github.com/alexia-23/go_final_project/pkg/logger"
	"github.com/alexia-23/go_final_project/pkg/utils"
	"net/http"
)

func ListTasks(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	tasks, err := db.SelectTasks()
	if err != nil {
		utils.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("tasks selected: ", len(tasks))
	resp := domain.TaskListResponse{Tasks: tasks}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
