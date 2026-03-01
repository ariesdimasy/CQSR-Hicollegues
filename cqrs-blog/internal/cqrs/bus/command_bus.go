package bus

import (
	"fmt"
	"reflect"
)

// CommandHandler is the interface that all command handlers must implement
type CommandHandler interface{}

// CommandBus dispatches commands to their registered handlers
type CommandBus struct {
	handlers map[reflect.Type]interface{}
}

// NewCommandBus creates a new CommandBus
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[reflect.Type]interface{}),
	}
}

// RegisterHandler registers a handler for a specific command type
func (cb *CommandBus) RegisterHandler(commandType interface{}, handler interface{}) {
	t := reflect.TypeOf(commandType)
	cb.handlers[t] = handler
}

// GetHandler retrieves the handler for a specific command type
func (cb *CommandBus) GetHandler(commandType interface{}) (interface{}, error) {
	t := reflect.TypeOf(commandType)
	handler, exists := cb.handlers[t]
	if !exists {
		return nil, fmt.Errorf("no handler registered for command type: %s", t.Name())
	}
	return handler, nil
}
