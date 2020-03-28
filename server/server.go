package server

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathan-osman/guessing-game/ui"
	"go.uber.org/zap"
)

// Server manages access to the game and hosts the front-end.
type Server struct {
	listener net.Listener
	logger   *zap.Logger
	stopped  chan bool
}

// New creates a new server instance.
func New(cfg *Config) (*Server, error) {
	l, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return nil, err
	}
	var (
		r = mux.NewRouter()
		s = &Server{
			listener: l,
			logger:   cfg.Logger.Named("server"),
			stopped:  make(chan bool),
		}
		server = http.Server{
			Handler: r,
		}
	)
	r.PathPrefix("/").Handler(http.FileServer(ui.Assets))
	go func() {
		defer close(s.stopped)
		defer s.logger.Info("server stopped")
		s.logger.Info("server started")
		if err := server.Serve(l); err != nil {
			s.logger.Error(err.Error())
		}
	}()
	return s, nil
}

// Close shuts down the server.
func (s *Server) Close() {
	s.listener.Close()
	<-s.stopped
}
