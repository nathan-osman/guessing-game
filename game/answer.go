package game

// Answer represents a player's answer to the question.
type Answer struct {
	Text        string `json:"text"`
	PlayerIndex int    `json:"-"`
}
