package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexia-23/go_final_project/pkg/logger"
	"github.com/alexia-23/go_final_project/pkg/utils"
)

func NextDate(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	params := r.URL.Query()

	now := params.Get("now")

	t, err := time.Parse("20060102", now)
	if err != nil {
		http.Error(w, "Now is required", http.StatusBadRequest)
		return
	}

	date := params.Get("date")
	repeat := params.Get("repeat")

	next, err := utils.NextDate(t, date, repeat)
	if err != nil {
		http.Error(w, "Failed to calculate next date", http.StatusInternalServerError)
		return
	}
	log.Println("next: ", next)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err = fmt.Fprint(w, next)
	if err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}
