package game

// Player represents a participant in the game.
type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Out   bool   `json:"out"`
}
