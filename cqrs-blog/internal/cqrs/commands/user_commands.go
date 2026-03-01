package commands

// CreateUserCommand represents a command to create a new user
type CreateUserCommand struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   uint   `json:"role_id" validate:"required"`
}

// UpdateUserCommand represents a command to update an existing user
type UpdateUserCommand struct {
	ID     uint   `json:"-"`
	Name   string `json:"name" validate:"omitempty,min=3,max=100"`
	Email  string `json:"email" validate:"omitempty,email"`
	RoleID uint   `json:"role_id"`
}

// DeleteUserCommand represents a command to delete a user
type DeleteUserCommand struct {
	ID uint
}
