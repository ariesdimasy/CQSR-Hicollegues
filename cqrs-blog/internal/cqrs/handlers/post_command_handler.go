package handlers

import (
	"cqrs-blog/internal/cqrs/commands"
	"cqrs-blog/internal/domain/post"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// PostCommandHandler handles all post-related commands
type PostCommandHandler struct {
	db       *gorm.DB
	validate *validator.Validate
}

// NewPostCommandHandler creates a new PostCommandHandler
func NewPostCommandHandler(db *gorm.DB, validate *validator.Validate) *PostCommandHandler {
	return &PostCommandHandler{
		db:       db,
		validate: validate,
	}
}

// HandleCreate handles the CreatePostCommand
func (h *PostCommandHandler) HandleCreate(cmd commands.CreatePostCommand) (*post.PostResponse, error) {
	if err := h.validate.Struct(cmd); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Verify user exists
	var count int64
	if err := h.db.Table("users").Where("id = ? AND deleted_at IS NULL", cmd.UserID).Count(&count).Error; err != nil {
		return nil, fmt.Errorf("failed to verify user: %w", err)
	}
	if count == 0 {
		return nil, errors.New("user not found")
	}

	newPost := &post.Post{
		Title:   cmd.Title,
		Content: cmd.Content,
		UserID:  cmd.UserID,
	}

	if err := h.db.Create(newPost).Error; err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return post.ToPostResponse(newPost), nil
}

// HandleUpdate handles the UpdatePostCommand
func (h *PostCommandHandler) HandleUpdate(cmd commands.UpdatePostCommand) (*post.PostResponse, error) {
	if err := h.validate.Struct(cmd); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	var p post.Post
	if err := h.db.First(&p, cmd.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	if cmd.Title != "" {
		p.Title = cmd.Title
	}
	if cmd.Content != "" {
		p.Content = cmd.Content
	}

	if err := h.db.Save(&p).Error; err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	return post.ToPostResponse(&p), nil
}

// HandleDelete handles the DeletePostCommand
func (h *PostCommandHandler) HandleDelete(cmd commands.DeletePostCommand) error {
	result := h.db.Delete(&post.Post{}, cmd.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found")
	}
	return nil
}
