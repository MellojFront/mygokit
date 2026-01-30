package web

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	Addr    string
	Handler http.Handler
}

func New(addr string, handler http.Handler) *Server {

	return &Server{
		Addr:    addr,
		Handler: handler,
	}
}

func (s *Server) Run() error {
	srv := http.Server{
		Addr:         s.Addr,
		Handler:      s.Handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("🔥 сервер запущен на http://localhost%s", s.Addr)
	return srv.ListenAndServe()
}
