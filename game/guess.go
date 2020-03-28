package game

// Guess represents a player's guess.
type Guess struct {
	PlayerIndex      int  `json:"player_index"`
	GuessPlayerIndex int  `json:"guess_player_index"`
	GuessAnswerIndex int  `json:"guess_answer_index"`
	Good             bool `json:"good"`
}
