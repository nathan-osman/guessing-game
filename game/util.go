package game

// playerIndexValid determines if the provided player index is valid.
func (g *Game) playerIndexValid(index int) bool {
	return index >= 0 && index < len(g.Players)
}

// answerIndexValid determines if the provided answer index is valid.
func (g *Game) answerIndexValid(index int) bool {
	return index >= 0 && index < len(g.Answers)
}

// answerExists checks to see if the provided answer already exists.
func (g *Game) answerExists(text string) bool {
	for _, a := range g.Answers {
		if a.Text == text {
			return true
		}
	}
	return false
}

// findAnswer searches for an answer by the provided player.
func (g *Game) findAnswer(playerIndex int) (int, *Answer) {
	for i, a := range g.Answers {
		if a.PlayerIndex == playerIndex {
			return i, a
		}
	}
	return IndexNone, nil
}

// remainingPlayers returns the number of players remaining in the game.
func (g *Game) remainingPlayers() int {
	remainingPlayers := len(g.Players)
	for _, p := range g.Players {
		if p.Out {
			remainingPlayers--
		}
	}
	return remainingPlayers
}

// Advance the current guesser to the next player who is not out.
func (g *Game) advanceGuesser() {
	for {
		g.GuesserIndex++
		if g.GuesserIndex >= len(g.Players) {
			g.GuesserIndex = 0
		}
		if !g.Players[g.GuesserIndex].Out {
			return
		}
	}
}

// wrap invokes the specified function with the mutex locked.
func (g *Game) wrap(state string, fn func() error) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if g.State != state {
		return errInvalidState
	}
	return fn()
}
