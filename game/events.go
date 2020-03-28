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

type playerAddedEvent struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
}

type playerRemovedEvent struct {
	GUID string `json:"guid"`
}

type questionAskedEvent struct {
	Question string `json:"question"`
}

type answersReceivedEvent struct {
	Answers map[string]*Answer `json:"answers"`
}

type guessMadeEvent struct {
	Guess *Guess `json:"guess"`
}

// sendEvent sends the specified event on the event channel.
func (g *Game) sendEvent(eventType string, data interface{}) {
	g.eventChan <- &baseEvent{
		Type: eventType,
		Data: data,
	}
}
