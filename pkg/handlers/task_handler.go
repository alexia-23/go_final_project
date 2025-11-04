package handlers

import (
	"net/http"
)

func HandleTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTask(w, r)
	case http.MethodPost:
		CreateTask(w, r)
	case http.MethodPut:
		PutTask(w, r)
	case http.MethodDelete:
		DeleteTask(w, r)
	default:
		WriteError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
