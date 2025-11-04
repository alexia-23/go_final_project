package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alexia-23/go_final_project/pkg/db"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	raw := params.Get("id")
	id, err := strconv.Atoi(raw)
	if err != nil {
		WriteError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	_, err = db.DeleteTaskById(id)

	if err == sql.ErrNoRows {
		WriteError(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]any{}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
