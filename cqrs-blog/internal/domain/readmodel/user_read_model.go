package readmodel

import (
	"cqrs-blog/internal/domain/user"
	"time"
)

// UserReadModel is the MongoDB read model for User
type UserReadModel struct {
	ID        uint      `bson:"_id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	RoleID    uint      `bson:"role_id" json:"role_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// ToUserResponse converts a UserReadModel to UserResponse
func (u *UserReadModel) ToUserResponse() *user.UserResponse {
	return &user.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		RoleID:    u.RoleID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// ToUserResponses converts a slice of UserReadModel to UserResponse slice
func ToUserResponses(users []UserReadModel) []user.UserResponse {
	responses := make([]user.UserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, *u.ToUserResponse())
	}
	return responses
}

// NewUserReadModel creates a UserReadModel from a domain User
func NewUserReadModel(u *user.User) *UserReadModel {
	return &UserReadModel{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		RoleID:    u.RoleID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
