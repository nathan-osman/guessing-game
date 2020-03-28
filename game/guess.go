package game

// Guess represents a player's guess.
type Guess struct {
	PlayerGUID      string `json:"player_guid"`
	GuessPlayerGUID string `json:"guess_player_guid"`
	GuessAnswerGUID string `json:"guess_answer_guid"`
	Good            bool   `json:"good"`
}
