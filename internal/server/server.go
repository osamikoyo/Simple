package server

import (
	"context"
	"net/http"

	"github.com/osamikoyo/simple/pkg/logger"
)

type Server struct{
	*http.Server
	logger *logger.Logger
}

func Init() *Server {
	return &Server{
		Server: &http.Server{
			Addr: "localhost:8080",
		},
		logger: logger.Init(),
	}
}

func (s *Server) Run() error {
	s.logger.Infof("starting server on port: %s", s.Addr)

	return s.ListenAndServe()
}

func (s *Server) SetHandler(mux *http.ServeMux) {
	s.logger.Info("set handler")

	s.Handler = mux
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}