package game

import (
	"encoding/json"
)

type gameState struct {
	GUID                string             `json:"guid"`
	Name                string             `json:"name"`
	State               string             `json:"state"`
	PlayerSequence      []string           `json:"player_sequence"`
	Players             map[string]*Player `json:"players"`
	AnswersByAnswerGUID map[string]*Answer `json:"answers"`
	AnswersByPlayerGUID map[string]*Answer `json:"-"`
	LastGuess           *Guess             `json:"last_guess"`
	Question            string             `json:"question"`
	GuesserIndex        int                `json:"guesser_index"`
}

// State encodes the game's current state as a JSON document.
func (g *Game) State() ([]byte, error) {
	var (
		data []byte
		err  = g.wrap(stateAny, func() error {
			b, err := json.Marshal(&g.state)
			if err != nil {
				return err
			}
			data = b
			return nil
		})
	)
	return data, err
}
