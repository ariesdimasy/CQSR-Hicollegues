package events

import "time"

type EventType string

const (
	UserCreated EventType = "user.created"
	UserUpdated EventType = "user.updated"
	UserDeleted EventType = "user.deleted"

	PostCreated EventType = "post.created"
	PostUpdated EventType = "post.updated"
	PostDeleted EventType = "post.deleted"

	RoleCreated EventType = "role.created"
	RoleUpdated EventType = "role.updated"
	RoleDeleted EventType = "role.deleted"
)

type Event struct {
	Type      EventType   `json:"type"`
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
}

// User Events
type UserCreatedEvent struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	RoleID    uint      `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUpdatedEvent struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	RoleID    uint      `json:"role_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDeletedEvent struct {
	ID        uint      `json:"id"`
	DeletedAt time.Time `json:"deleted_at"`
}

// Post Events
type PostCreatedEvent struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PostUpdatedEvent struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostDeletedEvent struct {
	ID        uint      `json:"id"`
	DeletedAt time.Time `json:"deleted_at"`
}

// Role Events
type RoleCreatedEvent struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type RoleUpdatedEvent struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleDeletedEvent struct {
	ID        uint      `json:"id"`
	DeletedAt time.Time `json:"deleted_at"`
}
