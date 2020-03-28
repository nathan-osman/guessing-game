package game

// Answer represents a player's answer to the question.
type Answer struct {
	GUID       string `json:"-"`
	Text       string `json:"text"`
	PlayerGUID string `json:"-"`
}
