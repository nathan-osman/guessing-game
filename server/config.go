package server

import (
	"go.uber.org/zap"
)

// Config stores the configuration for the server.
type Config struct {
	Addr   string
	Debug  bool
	Logger *zap.Logger
}
