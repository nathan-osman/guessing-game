package game

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

const (
	stateAny                = "*"
	stateWaitingForPlayers  = "waiting-for-players"
	stateWaitingForQuestion = "waiting-for-question"
	stateWaitingForAnswers  = "waiting-for-answers"
	stateWaitingForGuess    = "waiting-for-guess"
	stateWaitingForRestart  = "waiting-for-restart"
)

var (
	errInvalidAction    = errors.New("invalid action")
	errInvalidParameter = errors.New("invalid parameter")
	errDuplicateAnswer  = errors.New("duplicate answer")
)

// Game maintains game state during gameplay. Great care has been taken to make
// sure that the methods for this type are thread-safe since they will be
// invoked from many different goroutines.
type Game struct {
	mutex     sync.Mutex
	state     gameState
	eventChan chan<- string
}

// wrap invokes the specified function with the mutex locked and confirms that
// the game is in the specified state.
func (g *Game) wrap(state string, fn func() error) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if state != stateAny && g.state.State != state {
		return errInvalidAction
	}
	return fn()
}

// init resets the game to its initial state with the exception of player and
// player sequence data.
func (g *Game) init() {
	g.state.State = stateWaitingForPlayers
	if len(g.state.PlayerSequence) >= 2 {
		g.state.PlayerSequence = append(
			g.state.PlayerSequence[:1],
			g.state.PlayerSequence[0],
		)
	}
	if g.state.Players == nil {
		g.state.Players = map[string]*Player{}
	}
	g.state.AnswersByAnswerGUID = map[string]*Answer{}
	g.state.AnswersByPlayerGUID = map[string]*Answer{}
	g.state.LastGuess = nil
}

// New creates and initializes a new game.
func New(cfg *Config) *Game {
	g := &Game{
		state: gameState{
			GUID: uuid.Must(uuid.NewRandom()).String(),
			Name: cfg.Name,
		},
		eventChan: cfg.EventChan,
	}
	g.init()
	return g
}

// Add adds a new player to the game and returns their GUID. If the current
// asker is not set, then the user is assigned that position.
func (g *Game) Add(name string) (string, error) {
	var (
		playerGUID string
		err        = g.wrap(stateWaitingForPlayers, func() error {
			playerGUID = uuid.Must(uuid.NewRandom()).String()
			p := &Player{
				Name: name,
			}
			g.state.PlayerSequence = append(g.state.PlayerSequence, playerGUID)
			g.state.Players[playerGUID] = p
			g.sendEvent(eventPlayerAdded, &playerAddedEvent{
				GUID: playerGUID,
				Name: name,
			})
			return nil
		})
	)
	return playerGUID, err
}

// Remove removes the specified player from the game.
func (g *Game) Remove(playerGUID string) error {
	return g.wrap(stateWaitingForPlayers, func() error {
		for i, p := range g.state.PlayerSequence {
			if p == playerGUID {
				g.state.PlayerSequence[i] = g.state.PlayerSequence[len(g.state.PlayerSequence)-1]
				g.state.PlayerSequence = g.state.PlayerSequence[:len(g.state.PlayerSequence)-1]
				break
			}
		}
		delete(g.state.Players, playerGUID)
		g.sendEvent(eventPlayerRemoved, &playerRemovedEvent{
			GUID: playerGUID,
		})
		return nil
	})
}

// Start begins the game. Only the first player may do this and only when there
// are at least three players in the game.
func (g *Game) Start(playerGUID string) error {
	return g.wrap(stateWaitingForPlayers, func() error {
		if playerGUID != g.state.PlayerSequence[0] || len(g.state.Players) < 3 {
			return errInvalidAction
		}
		g.state.State = stateWaitingForQuestion
		g.sendEvent(eventGameStarted, nil)
		return nil
	})
}

// Ask submits the question for the game. Only the first player may do this.
func (g *Game) Ask(playerGUID, question string) error {
	return g.wrap(stateWaitingForQuestion, func() error {
		if playerGUID != g.state.PlayerSequence[0] {
			return errInvalidAction
		}
		g.state.State = stateWaitingForAnswers
		g.sendEvent(eventQuestionAsked, &questionAskedEvent{
			Question: question,
		})
		return nil
	})
}

// Answer submits an answer for the game. The answer cannot be a duplicate of
// any existing answers. If a player has already submitted an answer, it is
// modified.
func (g *Game) Answer(playerGUID, text string) error {
	return g.wrap(stateWaitingForAnswers, func() error {
		for _, a := range g.state.AnswersByAnswerGUID {
			if a.Text == text {
				return errDuplicateAnswer
			}
		}
		a := g.state.AnswersByPlayerGUID[playerGUID]
		if a == nil {
			answerGUID := uuid.Must(uuid.NewRandom()).String()
			a = &Answer{
				GUID:       answerGUID,
				PlayerGUID: playerGUID,
			}
			g.state.AnswersByAnswerGUID[answerGUID] = a
			g.state.AnswersByPlayerGUID[playerGUID] = a
		}
		a.Text = text
		if len(g.state.AnswersByAnswerGUID) == len(g.state.Players) {
			g.state.State = stateWaitingForGuess
			g.state.GuesserIndex = 1
			g.sendEvent(eventAnswersReceived, &answersReceivedEvent{
				Answers: g.state.AnswersByAnswerGUID,
			})
		}
		return nil
	})
}

// Guess registers a guess for a particular player. If the guess was
// successful, the player gains a point and continues guessing. If not, the
// next player takes their turn. If at any point only two players remain,
// the game is over and the game is reset (except for scores, which are
// accumulated).
func (g *Game) Guess(playerGUID, guessPlayerGUID, guessAnswerGUID string) error {
	return g.wrap(stateWaitingForGuess, func() error {
		if playerGUID != g.state.PlayerSequence[g.state.GuesserIndex] {
			return errInvalidAction
		}
		if _, ok := g.state.Players[guessPlayerGUID]; !ok {
			return errInvalidParameter
		}
		if _, ok := g.state.AnswersByAnswerGUID[guessAnswerGUID]; !ok {
			return errInvalidParameter
		}
		var (
			good  = g.state.AnswersByPlayerGUID[guessPlayerGUID].GUID == guessAnswerGUID
			guess = &Guess{
				PlayerGUID:      playerGUID,
				GuessPlayerGUID: guessPlayerGUID,
				GuessAnswerGUID: guessAnswerGUID,
				Good:            good,
			}
		)
		g.state.LastGuess = guess
		g.sendEvent(eventGuessMade, &guessMadeEvent{Guess: guess})
		if good {
			g.state.Players[playerGUID].Score++
			g.state.Players[guessPlayerGUID].Out = true
			remainingPlayers := len(g.state.Players)
			for _, p := range g.state.Players {
				if p.Out {
					remainingPlayers--
				}
			}
			if remainingPlayers <= 2 {
				g.state.State = stateWaitingForRestart
				g.sendEvent(eventGameFinished, nil)
			}
		} else {
			for {
				g.state.GuesserIndex++
				if g.state.GuesserIndex >= len(g.state.PlayerSequence) {
					g.state.GuesserIndex = 0
				}
				if !g.state.Players[g.state.PlayerSequence[g.state.GuesserIndex]].Out {
					break
				}
			}
		}
		return nil
	})
}

// Restart begins the game again, allowing new players to join and existing
// players to leave. Only the current asker may restart the game.
func (g *Game) Restart(playerGUID string) error {
	return g.wrap(stateWaitingForRestart, func() error {
		if playerGUID != g.state.PlayerSequence[0] {
			return errInvalidAction
		}
		g.init()
		g.sendEvent(eventGameRestarted, nil)
		return nil
	})
}
