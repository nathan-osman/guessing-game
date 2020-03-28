package game

import (
	"errors"
	"sync"
)

const (
	// StateLobby - waiting for players to leave or join.
	StateLobby = "lobby"
	// StateWaitingForQuestion - waiting for the question to be selected.
	StateWaitingForQuestion = "waiting-for-question"
	// StateWaitingForAnswers - waiting for players to submit their answers.
	StateWaitingForAnswers = "waiting-for-answers"
	// StateGuessing - waiting for a player to make a guess.
	StateGuessing = "guessing"
	// StateFinished - the game has completed.
	StateFinished = "finished"

	// IndexNone indicates a non-existent item.
	IndexNone = -1
)

var (
	errInvalidState = errors.New("you can't do that right now")
	errInvalidIndex = errors.New("invalid index")
	errUnauthorized = errors.New("you are not permitted to do that")
	errDuplicate    = errors.New("another player has already submitted this")
	errInvalidGuess = errors.New("your guess is invalid")
)

// Game maintains game state during gameplay. Great care has been taken to make
// sure that the methods for this type are thread-safe since they will be
// invoked from many different goroutines.
type Game struct {
	mutex sync.Mutex

	Name         string    `json:"name"`
	State        string    `json:"state"`
	Players      []*Player `json:"players"`
	Answers      []*Answer `json:"answers"`
	Guesses      []*Guess  `json:"guesses"`
	Question     string    `json:"question"`
	AskerIndex   int       `json:"asker_index"`
	GuesserIndex int       `json:"guesser_index"`
}

// Add adds a new player to the game and returns their index.
func (g *Game) Add(name string) (int, error) {
	var (
		playerIndex int
		err         = g.wrap(StateLobby, func() error {
			playerIndex = len(g.Players)
			g.Players = append(g.Players, &Player{Name: name})
			return nil
		})
	)
	return playerIndex, err
}

// Start begins the game. Only the current asker may do this.
func (g *Game) Start(playerIndex int) error {
	return g.wrap(StateLobby, func() error {
		if playerIndex != g.AskerIndex {
			return errUnauthorized
		}
		g.State = StateWaitingForQuestion
		return nil
	})
}

// Ask submits a question for the game.
func (g *Game) Ask(playerIndex int, question string) error {
	return g.wrap(StateWaitingForQuestion, func() error {
		if playerIndex != g.AskerIndex {
			return errUnauthorized
		}
		g.Question = question
		g.State = StateWaitingForAnswers
		return nil
	})
}

// Answer submits an answer for the game. The answer cannot be a duplicate of
// any existing answers
func (g *Game) Answer(playerIndex int, text string) error {
	return g.wrap(StateWaitingForAnswers, func() error {
		if g.answerExists(text) {
			return errDuplicate
		}
		i, a := g.findAnswer(playerIndex)
		if i == IndexNone {
			i = len(g.Answers)
			a = &Answer{
				PlayerIndex: playerIndex,
			}
			g.Answers = append(g.Answers, a)
		}
		a.Text = text
		if len(g.Answers) == len(g.Players) {
			g.State = StateGuessing
		}
		return nil
	})
}

// Guess registers a guess for a particular player.
func (g *Game) Guess(playerIndex, guessPlayerIndex, guessAnswerIndex int) error {
	return g.wrap(StateGuessing, func() error {
		if playerIndex != g.GuesserIndex {
			return errUnauthorized
		}
		if !g.playerIndexValid(guessPlayerIndex) ||
			!g.answerIndexValid(guessAnswerIndex) {
			return errInvalidIndex
		}
		if g.Players[guessPlayerIndex].Out {
			return errInvalidGuess
		}
		good := false
		i, _ := g.findAnswer(guessPlayerIndex)
		if i == guessAnswerIndex {
			good = true
		}
		guess := &Guess{
			PlayerIndex:      playerIndex,
			GuessPlayerIndex: guessPlayerIndex,
			GuessAnswerIndex: guessAnswerIndex,
			Good:             good,
		}
		g.Guesses = append(g.Guesses, guess)
		if good {
			g.Players[playerIndex].Score++
			g.Players[guessPlayerIndex].Out = true
			if g.remainingPlayers() == len(g.Players)-1 {
				g.State = StateFinished
			}
		} else {
			g.advanceGuesser()
		}
		return nil
	})
}
