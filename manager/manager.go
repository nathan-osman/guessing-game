package manager

import (
	"encoding/json"
	"sync"

	"github.com/nathan-osman/guessing-game/game"
	"go.uber.org/zap"
)

// Manager wraps a game instance and manages client connections.
type Manager struct {
	mutex       sync.Mutex
	game        *game.Game
	logger      *zap.Logger
	clients     map[string]*managerClient
	eventChan   <-chan interface{}
	stopChan    chan bool
	stoppedChan chan bool
}

func (m *Manager) sendToClients(v interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	b, _ := json.Marshal(v)
	s := string(b)
	for _, c := range m.clients {
		c.Send(s)
	}
}

func (m *Manager) run() {
	defer close(m.stopChan)
	defer m.logger.Info("manager stopped")
	m.logger.Info("manager started")
	for {
		select {
		case v := <-m.eventChan:
			m.sendToClients(v)
		case <-m.stopChan:
			return
		}
	}
}

// New creates a new game and returns its GUID.
func New(cfg *Config) *Manager {
	var (
		eventChan = make(chan interface{})
		m         = &Manager{
			game: game.New(&game.Config{
				Name:      cfg.Name,
				EventChan: eventChan,
			}),
			logger:      cfg.Logger.Named("manager"),
			clients:     map[string]*managerClient{},
			eventChan:   eventChan,
			stopChan:    make(chan bool),
			stoppedChan: make(chan bool),
		}
	)
	go m.run()
	return m
}

// Name returns the name of the game.
func (m *Manager) Name() string {
	return m.game.Name()
}

// Connect is invoked when a new connection is received.
func (m *Manager) Connect() {
	//...
}

// Close shuts down the manager.
func (m *Manager) Close() {
	close(m.stopChan)
	<-m.stoppedChan
}
