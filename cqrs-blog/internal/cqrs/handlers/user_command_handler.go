package handlers

import (
	"cqrs-blog/internal/cqrs/commands"
	"cqrs-blog/internal/domain/user"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserCommandHandler handles all user-related commands
type UserCommandHandler struct {
	db       *gorm.DB
	validate *validator.Validate
}

// NewUserCommandHandler creates a new UserCommandHandler
func NewUserCommandHandler(db *gorm.DB, validate *validator.Validate) *UserCommandHandler {
	return &UserCommandHandler{
		db:       db,
		validate: validate,
	}
}

// HandleCreate handles the CreateUserCommand
func (h *UserCommandHandler) HandleCreate(cmd commands.CreateUserCommand) (*user.UserResponse, error) {
	if err := h.validate.Struct(cmd); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := &user.User{
		Name:     cmd.Name,
		Email:    cmd.Email,
		Password: string(hashedPassword),
		RoleID:   cmd.RoleID,
	}

	if err := h.db.Create(newUser).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user.ToUserResponse(newUser), nil
}

// HandleUpdate handles the UpdateUserCommand
func (h *UserCommandHandler) HandleUpdate(cmd commands.UpdateUserCommand) (*user.UserResponse, error) {
	if err := h.validate.Struct(cmd); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	var u user.User
	if err := h.db.First(&u, cmd.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if cmd.Name != "" {
		u.Name = cmd.Name
	}
	if cmd.Email != "" {
		u.Email = cmd.Email
	}
	if cmd.RoleID != 0 {
		u.RoleID = cmd.RoleID
	}

	if err := h.db.Save(&u).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user.ToUserResponse(&u), nil
}

// HandleDelete handles the DeleteUserCommand
func (h *UserCommandHandler) HandleDelete(cmd commands.DeleteUserCommand) error {
	result := h.db.Delete(&user.User{}, cmd.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
