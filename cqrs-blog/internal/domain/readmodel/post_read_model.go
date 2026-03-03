package readmodel

import (
	"cqrs-blog/internal/domain/post"
	"time"
)

// PostReadModel is the MongoDB read model for Post
type PostReadModel struct {
	ID        uint      `bson:"_id" json:"id"`
	Title     string    `bson:"title" json:"title"`
	Content   string    `bson:"content" json:"content"`
	UserID    uint      `bson:"user_id" json:"user_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// ToPostResponse converts a PostReadModel to PostResponse
func (p *PostReadModel) ToPostResponse() *post.PostResponse {
	return &post.PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToPostResponses converts a slice of PostReadModel to PostResponse slice
func ToPostResponses(posts []PostReadModel) []post.PostResponse {
	responses := make([]post.PostResponse, 0, len(posts))
	for _, p := range posts {
		responses = append(responses, *p.ToPostResponse())
	}
	return responses
}

// NewPostReadModel creates a PostReadModel from a domain Post
func NewPostReadModel(p *post.Post) *PostReadModel {
	return &PostReadModel{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
