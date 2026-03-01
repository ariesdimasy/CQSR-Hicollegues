package services

import (
	"errors"
	"event-driven-blog/internal/application/events"
	"event-driven-blog/internal/domain/user"
	"event-driven-blog/internal/infrastructure/eventbus"

	"event-driven-blog/internal/infrastructure/validator"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db       *gorm.DB
	eventBus *eventbus.EventBus
	validate *validator.Validator
}

func NewUserService(db *gorm.DB, eventBus *eventbus.EventBus, validate *validator.Validator) *UserService {
	return &UserService{
		db:       db,
		eventBus: eventBus,
		validate: validate,
	}
}

func (s *UserService) Create(req *user.CreateUserRequest) (*user.UserResponse, error) {
	// Validate request
	if err := s.validate.Validate(req); err != nil {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	newUser := &user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		RoleID:   req.RoleID,
	}

	if err := s.db.Create(newUser).Error; err != nil {
		return nil, err
	}

	// Publish event
	s.eventBus.Publish(events.Event{
		Type:      events.UserCreated,
		Timestamp: time.Now(),
		Data: events.UserCreatedEvent{
			ID:        newUser.ID,
			Name:      newUser.Name,
			Email:     newUser.Email,
			RoleID:    newUser.RoleID,
			CreatedAt: newUser.CreatedAt,
		},
	})

	return &user.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		RoleID:    newUser.RoleID,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}, nil
}

func (s *UserService) GetByID(id uint) (*user.UserResponse, error) {
	var u user.User
	if err := s.db.First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		RoleID:    u.RoleID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (s *UserService) GetAll() ([]user.UserResponse, error) {
	var users []user.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}

	var responses []user.UserResponse
	for _, u := range users {
		responses = append(responses, user.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			RoleID:    u.RoleID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return responses, nil
}

func (s *UserService) Update(id uint, req *user.UpdateUserRequest) (*user.UserResponse, error) {
	var u user.User
	if err := s.db.First(&u, id).Error; err != nil {
		return nil, err
	}

	if req.Name != "" {
		u.Name = req.Name
	}
	if req.Email != "" {
		u.Email = req.Email
	}
	if req.RoleID != 0 {
		u.RoleID = req.RoleID
	}

	if err := s.db.Save(&u).Error; err != nil {
		return nil, err
	}

	// Publish event
	s.eventBus.Publish(events.Event{
		Type:      events.UserUpdated,
		Timestamp: time.Now(),
		Data: events.UserUpdatedEvent{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			RoleID:    u.RoleID,
			UpdatedAt: u.UpdatedAt,
		},
	})

	return &user.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		RoleID:    u.RoleID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (s *UserService) Delete(id uint) error {
	result := s.db.Delete(&user.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	// Publish event
	s.eventBus.Publish(events.Event{
		Type:      events.UserDeleted,
		Timestamp: time.Now(),
		Data: events.UserDeletedEvent{
			ID:        id,
			DeletedAt: time.Now(),
		},
	})

	return nil
}
