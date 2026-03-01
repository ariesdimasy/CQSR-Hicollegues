package bus

import (
	"fmt"
	"reflect"
)

// QueryHandler is the interface that all query handlers must implement
type QueryHandler interface{}

// QueryBus dispatches queries to their registered handlers
type QueryBus struct {
	handlers map[reflect.Type]interface{}
}

// NewQueryBus creates a new QueryBus
func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[reflect.Type]interface{}),
	}
}

// RegisterHandler registers a handler for a specific query type
func (qb *QueryBus) RegisterHandler(queryType interface{}, handler interface{}) {
	t := reflect.TypeOf(queryType)
	qb.handlers[t] = handler
}

// GetHandler retrieves the handler for a specific query type
func (qb *QueryBus) GetHandler(queryType interface{}) (interface{}, error) {
	t := reflect.TypeOf(queryType)
	handler, exists := qb.handlers[t]
	if !exists {
		return nil, fmt.Errorf("no handler registered for query type: %s", t.Name())
	}
	return handler, nil
}
