package user

import (
	"time"

	"gorm.io/gorm"
)

// Entity
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name" validate:"required,min=3,max=100"`
	Email     string         `gorm:"size:100;uniqueIndex;not null" json:"email" validate:"required,email"`
	Password  string         `gorm:"size:255;not null" json:"-" validate:"required,min=6"`
	RoleID    uint           `gorm:"not null" json:"role_id" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Response DTO
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	RoleID    uint      `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToUserResponse(u *User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		RoleID:    u.RoleID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToUserResponses(users []User) []UserResponse {
	var responses []UserResponse
	for _, u := range users {
		responses = append(responses, *ToUserResponse(&u))
	}
	return responses
}
