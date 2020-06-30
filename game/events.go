package game

const (
	eventPlayerAdded     = "player-added"
	eventPlayerRemoved   = "player-removed"
	eventGameStarted     = "game-started"
	eventQuestionAsked   = "question-asked"
	eventAnswersReceived = "answers-received"
	eventGuessMade       = "guess-made"
	eventGameFinished    = "game-finished"
	eventGameRestarted   = "game-restarted"
)

type baseEvent struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type playerAddedData struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
}

type playerRemovedData struct {
	GUID string `json:"guid"`
}

type gameStartedData struct {
	PlayerSequence []string `json:"player_sequence"`
}

type questionAskedData struct {
	Question string `json:"question"`
}

type answersReceivedData struct {
	Answers map[string]*Answer `json:"answers"`
}

type guessMadeData struct {
	Guess *Guess `json:"guess"`
}
