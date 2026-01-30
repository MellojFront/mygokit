package server

import (
	"log"
	"net/http"
	"time"
)

// Server — конфигурациясервера
type Server struct {
	Addr      string
	StaticDir string
}

// New — конструктор для создания сервера
func New(addr string, staticDir string) *Server {
	// Если адрес пустой, ставим по умолчанию
	if addr == "" {
		addr = ":8080"
	}
	return &Server{
		Addr:      addr,
		StaticDir: staticDir,
	}
}

// Run — метод, который запускает сервер
func (s *Server) Run() error {
	mux := http.NewServeMux()

	// Настраиваем раздачу файлов
	// Если StaticDir не указан, не запускаем файловый сервер
	if s.StaticDir != "" {
		fileServer := http.FileServer(http.Dir(s.StaticDir))
		mux.Handle("/", fileServer)
	}

	srv := &http.Server{
		Addr:         s.Addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("🔥 сервер запущен на http://localhost%s", s.Addr)
	return srv.ListenAndServe()
}
