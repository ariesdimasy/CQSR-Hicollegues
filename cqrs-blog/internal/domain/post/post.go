package post

import (
	"time"

	"gorm.io/gorm"
)

// Entity
type Post struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `gorm:"size:200;not null" json:"title" validate:"required,min=3,max=200"`
	Content   string         `gorm:"type:text;not null" json:"content" validate:"required,min=10"`
	UserID    uint           `gorm:"not null;constraint:OnDelete:CASCADE" json:"user_id" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Response DTO
type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToPostResponse(p *Post) *PostResponse {
	return &PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ToPostResponses(posts []Post) []PostResponse {
	responses := make([]PostResponse, 0, len(posts))
	for _, p := range posts {
		responses = append(responses, *ToPostResponse(&p))
	}
	return responses
}
