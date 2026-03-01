package handlers

import (
	"cqrs-blog/internal/cqrs/queries"
	"cqrs-blog/internal/domain/user"
	"errors"

	"gorm.io/gorm"
)

// UserQueryHandler handles all user-related queries
type UserQueryHandler struct {
	db *gorm.DB
}

// NewUserQueryHandler creates a new UserQueryHandler
func NewUserQueryHandler(db *gorm.DB) *UserQueryHandler {
	return &UserQueryHandler{db: db}
}

// HandleGetByID handles the GetUserByIDQuery
func (h *UserQueryHandler) HandleGetByID(query queries.GetUserByIDQuery) (*user.UserResponse, error) {
	var u user.User
	if err := h.db.First(&u, query.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user.ToUserResponse(&u), nil
}

// HandleGetAll handles the GetAllUsersQuery
func (h *UserQueryHandler) HandleGetAll(query queries.GetAllUsersQuery) ([]user.UserResponse, error) {
	var users []user.User
	if err := h.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return user.ToUserResponses(users), nil
}
