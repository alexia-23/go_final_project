package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexia-23/go_final_project/pkg/storage"
)

type Server struct {
	Storage *storage.Storage
	Logger  *log.Logger
	Mux     *http.ServeMux
}

// statusRecorder нужен для перехвата кода ответа HTTP
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// loggingMiddleware — логирует запросы с временем выполнения и статусом
func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rec, r)

		s.Logger.Printf("[%s] %s %d %v",
			r.Method,
			r.URL.Path,
			rec.status,
			time.Since(start),
		)
	})
}

// New — конструктор, создаёт сервер и настраивает маршруты
func New(db *storage.Storage) *Server {
	s := &Server{
		Storage: db,
		Logger:  log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds),
		Mux:     http.NewServeMux(),
	}

	// Настройка маршрутов
	fs := http.FileServer(http.Dir("./web"))

	s.Mux.HandleFunc("/api/nextdate", s.HandleNextDate)
	s.Mux.HandleFunc("/api/task/done", s.PostTaskDone)
	s.Mux.HandleFunc("/api/task", s.HandleTask)
	s.Mux.HandleFunc("/api/tasks", s.ListTasks)
	s.Mux.Handle("/", fs)

	return s
}

// Run — запускает HTTP-сервер
func (s *Server) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7540"
	}

	s.Logger.Println("Сервер запущен на http://localhost:" + port)

	handler := s.loggingMiddleware(s.Mux)
	return http.ListenAndServe(":"+port, handler)
}
