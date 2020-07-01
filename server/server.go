package server

import (
	"net"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/nathan-osman/guessing-game/manager"
	"github.com/nathan-osman/guessing-game/ui"
	"go.uber.org/zap"
)

// Server manages access to the game and hosts the front-end.
type Server struct {
	mutex    sync.Mutex
	listener net.Listener
	logger   *zap.Logger
	managers map[string]*manager.Manager
	stopped  chan bool
}

// New creates a new server instance.
func New(cfg *Config) (*Server, error) {
	l, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return nil, err
	}
	var (
		r = chi.NewRouter()
		s = &Server{
			listener: l,
			logger:   cfg.Logger.Named("server"),
			managers: map[string]*manager.Manager{},
			stopped:  make(chan bool),
		}
		server = http.Server{
			Handler: r,
		}
	)
	r.Mount("/", http.FileServer(ui.Assets))
	r.Route("/api", func(r chi.Router) {
		r.Get("/games", s.apiGames)
		r.Post("/join", s.apiJoin)
		r.Post("/create", s.apiCreate)
	})
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
