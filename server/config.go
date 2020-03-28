package server

import (
	"go.uber.org/zap"
)

// Config stores the configuration for the server.
type Config struct {
	Addr   string
	Logger *zap.Logger
}
