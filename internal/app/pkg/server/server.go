package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/config"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/controller"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewServer(cfg *config.Config, handler *controller.Handler) *Server {
	return &Server{httpServer: &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Sever.Host, cfg.Sever.Port),
		Handler:      handler.Setup(),
		ReadTimeout:  time.Duration(cfg.Sever.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Sever.WriteTimeout) * time.Second,
	}}

}
