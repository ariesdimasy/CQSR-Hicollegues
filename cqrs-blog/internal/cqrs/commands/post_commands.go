package commands

// CreatePostCommand represents a command to create a new post
type CreatePostCommand struct {
	Title   string `json:"title" validate:"required,min=3,max=200"`
	Content string `json:"content" validate:"required,min=10"`
	UserID  uint   `json:"user_id" validate:"required"`
}

// UpdatePostCommand represents a command to update an existing post
type UpdatePostCommand struct {
	ID      uint   `json:"-"`
	Title   string `json:"title" validate:"omitempty,min=3,max=200"`
	Content string `json:"content" validate:"omitempty,min=10"`
}

// DeletePostCommand represents a command to delete a post
type DeletePostCommand struct {
	ID uint
}
