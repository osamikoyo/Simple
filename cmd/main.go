package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/osamikoyo/simple/internal/handler"
	"github.com/osamikoyo/simple/internal/server"
	"github.com/osamikoyo/simple/pkg/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	s := server.Init()
	logger := logger.Init()

	logger.Info("starting...")

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	s.SetHandler(mux)

	go func(){
		<-ctx.Done()
		s.Shutdown(ctx)
	}()

	if err := s.Run();err != nil && err != http.ErrServerClosed{
		logger.Error(err)
	}
}