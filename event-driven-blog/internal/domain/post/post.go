package post

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `gorm:"size:200;not null" json:"title" validate:"required,min=3,max=200"`
	Content   string         `gorm:"type:text;not null" json:"content" validate:"required,min=10"`
	UserID    uint           `gorm:"not null" json:"user_id" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreatePostRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=200"`
	Content string `json:"content" validate:"required,min=10"`
	UserID  uint   `json:"user_id" validate:"required"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" validate:"omitempty,min=3,max=200"`
	Content string `json:"content" validate:"omitempty,min=10"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
