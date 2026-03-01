package user

import (
	"time"

	"gorm.io/gorm"
)

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

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   uint   `json:"role_id" validate:"required"`
}

type UpdateUserRequest struct {
	Name   string `json:"name" validate:"omitempty,min=3,max=100"`
	Email  string `json:"email" validate:"omitempty,email"`
	RoleID uint   `json:"role_id"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	RoleID    uint      `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
