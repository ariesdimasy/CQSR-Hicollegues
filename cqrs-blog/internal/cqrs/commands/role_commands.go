package commands

// CreateRoleCommand represents a command to create a new role
type CreateRoleCommand struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}

// UpdateRoleCommand represents a command to update an existing role
type UpdateRoleCommand struct {
	ID   uint   `json:"-"`
	Name string `json:"name" validate:"omitempty,min=2,max=50"`
}

// DeleteRoleCommand represents a command to delete a role
type DeleteRoleCommand struct {
	ID uint
}
