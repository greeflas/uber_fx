package server

import (
	"context"
	"net"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type HTTPServer struct {
	srv      *http.Server
	listener net.Listener
	log      *logrus.Logger
}

func NewHTTPServer(p HTTPServerParams) *HTTPServer {
	return &HTTPServer{
		srv: &http.Server{
			Handler: p.Mux,
		},
		listener: p.Listener,
		log:      p.Log,
	}
}

type HTTPServerParams struct {
	fx.In

	Listener net.Listener
	Mux      *http.ServeMux
	Log      *logrus.Logger
}

func (s *HTTPServer) Start() error {
	go func() {
		if err := s.srv.Serve(s.listener); err != nil && err != http.ErrServerClosed {
			s.log.Errorf("Failed to serve listener: %v", err)
		}
	}()

	return nil
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func Run(lc fx.Lifecycle, log *logrus.Logger, srv *HTTPServer) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.WithFields(logrus.Fields{"addr": srv.listener.Addr().String()}).Info("Starting HTTP server")

			return srv.Start()
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
