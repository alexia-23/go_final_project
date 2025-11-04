package server

import (
	"net/http"
)

func (s *Server) HandleTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.GetTask(w, r)
	case http.MethodPost:
		s.CreateTask(w, r)
	case http.MethodPut:
		s.PutTask(w, r)
	case http.MethodDelete:
		s.DeleteTask(w, r)
	default:
		s.WriteError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
