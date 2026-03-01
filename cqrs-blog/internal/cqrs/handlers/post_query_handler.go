package handlers

import (
	"cqrs-blog/internal/cqrs/queries"
	"cqrs-blog/internal/domain/post"
	"errors"

	"gorm.io/gorm"
)

// PostQueryHandler handles all post-related queries
type PostQueryHandler struct {
	db *gorm.DB
}

// NewPostQueryHandler creates a new PostQueryHandler
func NewPostQueryHandler(db *gorm.DB) *PostQueryHandler {
	return &PostQueryHandler{db: db}
}

// HandleGetByID handles the GetPostByIDQuery
func (h *PostQueryHandler) HandleGetByID(query queries.GetPostByIDQuery) (*post.PostResponse, error) {
	var p post.Post
	if err := h.db.First(&p, query.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	return post.ToPostResponse(&p), nil
}

// HandleGetAll handles the GetAllPostsQuery
func (h *PostQueryHandler) HandleGetAll(query queries.GetAllPostsQuery) ([]post.PostResponse, error) {
	var posts []post.Post
	if err := h.db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return post.ToPostResponses(posts), nil
}

// HandleGetByUserID handles the GetPostsByUserIDQuery
func (h *PostQueryHandler) HandleGetByUserID(query queries.GetPostsByUserIDQuery) ([]post.PostResponse, error) {
	var posts []post.Post
	if err := h.db.Where("user_id = ?", query.UserID).Find(&posts).Error; err != nil {
		return nil, err
	}

	return post.ToPostResponses(posts), nil
}
