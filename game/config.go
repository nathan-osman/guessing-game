package game

// Config stores the configuration data for the game.
type Config struct {
	Name      string
	EventChan chan<- string
}
