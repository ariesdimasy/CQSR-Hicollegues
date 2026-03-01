package eventbus

import (
	"event-driven-blog/internal/application/events"
	"fmt"
	"sync"
)

type HandlerFunc func(event events.Event)

type EventBus struct {
	subscribers map[events.EventType][]HandlerFunc
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[events.EventType][]HandlerFunc),
	}
}

func (eb *EventBus) Subscribe(eventType events.EventType, handler HandlerFunc) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.subscribers[eventType] = append(eb.subscribers[eventType], handler)
}

func (eb *EventBus) Publish(event events.Event) {
	eb.mu.RLock()
	handlers, exists := eb.subscribers[event.Type]
	eb.mu.RUnlock()

	if exists {
		for _, handler := range handlers {
			go handler(event)
		}
	}
}

// Example handlers
func LogEventHandler(event events.Event) {
	fmt.Printf("[EVENT] Type: %s, Time: %s, Data: %+v\n",
		event.Type, event.Timestamp, event.Data)
}

func SendNotificationHandler(event events.Event) {
	switch event.Type {
	case events.UserCreated:
		fmt.Println("Sending welcome email to new user...")
	case events.PostCreated:
		fmt.Println("Notifying followers about new post...")
	}
}
