package manager

import (
	"go.uber.org/zap"
)

// Config stores the configuration for the manager.
type Config struct {
	Name   string
	Logger *zap.Logger
}
