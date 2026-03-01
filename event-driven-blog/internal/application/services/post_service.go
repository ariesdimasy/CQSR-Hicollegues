package services

import (
	"errors"
	"event-driven-blog/internal/application/events"
	"event-driven-blog/internal/domain/post"
	"event-driven-blog/internal/infrastructure/eventbus"
	"event-driven-blog/internal/infrastructure/validator"
	"time"

	"gorm.io/gorm"
)

type PostService struct {
	db       *gorm.DB
	eventBus *eventbus.EventBus
	validate *validator.Validator
}

func NewPostService(db *gorm.DB, eventBus *eventbus.EventBus, validate *validator.Validator) *PostService {
	return &PostService{
		db:       db,
		eventBus: eventBus,
		validate: validate,
	}
}

func (s *PostService) Create(req *post.CreatePostRequest) (*post.PostResponse, error) {
	if err := s.validate.Validate(req); err != nil {
		return nil, err
	}

	// Check if user exists
	var userExists bool
	s.db.Table("users").Select("count(*) > 0").Where("id = ?", req.UserID).Find(&userExists)
	if !userExists {
		return nil, errors.New("user not found")
	}

	newPost := &post.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := s.db.Create(newPost).Error; err != nil {
		return nil, err
	}

	// Publish event
	s.eventBus.Publish(events.Event{
		Type:      events.PostCreated,
		Timestamp: time.Now(),
		Data: events.PostCreatedEvent{
			ID:        newPost.ID,
			Title:     newPost.Title,
			Content:   newPost.Content,
			UserID:    newPost.UserID,
			CreatedAt: newPost.CreatedAt,
		},
	})

	return &post.PostResponse{
		ID:        newPost.ID,
		Title:     newPost.Title,
		Content:   newPost.Content,
		UserID:    newPost.UserID,
		CreatedAt: newPost.CreatedAt,
		UpdatedAt: newPost.UpdatedAt,
	}, nil
}

func (s *PostService) GetByID(id uint) (*post.PostResponse, error) {
	var p post.Post
	if err := s.db.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	return &post.PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (s *PostService) GetByUserID(userID uint) ([]post.PostResponse, error) {
	var posts []post.Post
	if err := s.db.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}

	var responses []post.PostResponse
	for _, p := range posts {
		responses = append(responses, post.PostResponse{
			ID:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			UserID:    p.UserID,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	return responses, nil
}

func (s *PostService) GetAll() ([]post.PostResponse, error) {
	var posts []post.Post
	if err := s.db.Find(&posts).Error; err != nil {
		return nil, err
	}

	var responses []post.PostResponse
	for _, p := range posts {
		responses = append(responses, post.PostResponse{
			ID:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			UserID:    p.UserID,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	return responses, nil
}

func (s *PostService) Update(id uint, req *post.UpdatePostRequest) (*post.PostResponse, error) {
	var p post.Post
	if err := s.db.First(&p, id).Error; err != nil {
		return nil, err
	}

	if req.Title != "" {
		p.Title = req.Title
	}
	if req.Content != "" {
		p.Content = req.Content
	}

	if err := s.db.Save(&p).Error; err != nil {
		return nil, err
	}

	// Publish event
	s.eventBus.Publish(events.Event{
		Type:      events.PostUpdated,
		Timestamp: time.Now(),
		Data: events.PostUpdatedEvent{
			ID:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			UserID:    p.UserID,
			UpdatedAt: p.UpdatedAt,
		},
	})

	return &post.PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (s *PostService) Delete(id uint) error {
	result := s.db.Delete(&post.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found")
	}

	// Publish event
	s.eventBus.Publish(events.Event{
		Type:      events.PostDeleted,
		Timestamp: time.Now(),
		Data: events.PostDeletedEvent{
			ID:        id,
			DeletedAt: time.Now(),
		},
	})

	return nil
}
