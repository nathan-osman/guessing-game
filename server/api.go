package server

import (
	"net/http"

	"github.com/go-chi/render"
)

type apiGame struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func (*apiGame) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) apiGames(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	games := []render.Renderer{}
	for uuid, m := range s.managers {
		games = append(games, &apiGame{
			UUID: uuid,
			Name: m.Name(),
		})
	}
	render.RenderList(w, r, games)
}

func (s *Server) apiJoin(w http.ResponseWriter, r *http.Request) {
	//...
}

func (s *Server) apiCreate(w http.ResponseWriter, r *http.Request) {
	//...
}
