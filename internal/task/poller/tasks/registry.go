// Package tasks provides poller task implementations.
package tasks

import (
	"go.uber.org/zap"

	"github.com/xelarion/go-layout/internal/task/poller"
)

// TaskHandler defines the interface that all poller task handlers must implement
type TaskHandler interface {
	Register(p *poller.Poller) error
}

// TaskHandlerConstructor defines a function type for creating task handlers
type TaskHandlerConstructor func(*zap.Logger) TaskHandler

// taskRegistry holds all registered task handlers
var taskRegistry = map[string]TaskHandlerConstructor{
	"metrics-collect": func(logger *zap.Logger) TaskHandler {
		return NewMetricsCollectHandler(logger)
	},
	// Add new task constructors here
}

// RegisterAll registers all poller tasks with the provided poller
func RegisterAll(p *poller.Poller, logger *zap.Logger) error {
	for taskName, constructorFn := range taskRegistry {
		taskHandler := constructorFn(logger)
		if err := taskHandler.Register(p); err != nil {
			logger.Error("Failed to register poller handler",
				zap.String("name", taskName),
				zap.Error(err))
			return err
		}
		logger.Info("Registered poller handler", zap.String("name", taskName))
	}
	return nil
}
