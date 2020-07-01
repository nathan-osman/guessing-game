package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nathan-osman/guessing-game/manager"
)

var (
	errEmptyName = errors.New("a valid name must be supplied for the game")
)

type apiGame struct {
	Name string `json:"name"`
}

func (g *apiGame) Bind(r *http.Request) error {
	if g.Name == "" {
		return errEmptyName
	}
	return nil
}

type apiGameReadOnly struct {
	*apiGame
	UUID string `json:"uuid"`
}

func (*apiGameReadOnly) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type apiError struct {
	Error string `json:"error"`
}

func (*apiError) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) apiGames(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	games := []render.Renderer{}
	for uuid, m := range s.managers {
		games = append(games, &apiGameReadOnly{
			apiGame: &apiGame{
				Name: m.Name(),
			},
			UUID: uuid,
		})
	}
	render.RenderList(w, r, games)
}

func (s *Server) apiJoin(w http.ResponseWriter, r *http.Request) {
	//...
}

func (s *Server) apiCreate(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	game := &apiGame{}
	if err := render.Bind(r, game); err != nil {
		render.Render(w, r, &apiError{
			Error: err.Error(),
		})
		return
	}
	gameUUID := uuid.Must(uuid.NewRandom()).String()
	s.managers[gameUUID] = manager.New(&manager.Config{
		Name:   game.Name,
		Logger: s.baseLogger,
	})
	render.Render(w, r, &apiGameReadOnly{
		apiGame: game,
		UUID:    gameUUID,
	})
}
