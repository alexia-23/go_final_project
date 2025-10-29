package server

import (
	"github.com/alexia-23/go_final_project/pkg/handlers"
	"github.com/alexia-23/go_final_project/pkg/logger"
	"log"
	"net/http"
	"os"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	log := logger.Get()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// обёртка для ResponseWriter, чтобы ловить статус
		rec := &statusRecorder{ResponseWriter: w, status: 200}

		next.ServeHTTP(rec, r)

		log.Printf("[%s] %s %d %v",
			r.Method,
			r.URL.Path,
			rec.status,
			time.Since(start),
		)
	})
}

func Init() error {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web"))
	mux.HandleFunc("/api/nextdate", handlers.NextDate)
	mux.HandleFunc("/api/task", handlers.CreateTask)
	mux.HandleFunc("/api/tasks", handlers.ListTasks)
	mux.Handle("/", fs)
	handler := loggingMiddleware(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7540"
	}

	log.Println("Сервер запущен на http://localhost:" + port)

	return http.ListenAndServe(":"+port, handler)
}
