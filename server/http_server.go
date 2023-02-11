package server

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
	"net/http"
)

type HTTPServer struct {
	srv      *http.Server
	listener net.Listener
	log      *zap.Logger
}

func NewHTTPServer(listener net.Listener, mux *http.ServeMux, log *zap.Logger) *HTTPServer {
	return &HTTPServer{
		srv: &http.Server{
			Handler: mux,
		},
		listener: listener,
		log:      log,
	}
}

func (s *HTTPServer) Start() error {
	go func() {
		if err := s.srv.Serve(s.listener); err != nil && err != http.ErrServerClosed {
			s.log.Error("Failed to serve listener", zap.Error(err))
		}
	}()

	return nil
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func Run(lc fx.Lifecycle, log *zap.Logger, srv *HTTPServer) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting HTTP server", zap.String("addr", srv.listener.Addr().String()))

			return srv.Start()
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
