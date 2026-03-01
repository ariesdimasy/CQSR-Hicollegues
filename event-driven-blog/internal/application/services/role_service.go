package services

import (
	"errors"
	"event-driven-blog/internal/application/events"
	"event-driven-blog/internal/domain/role"
	"event-driven-blog/internal/infrastructure/eventbus"
	"event-driven-blog/internal/infrastructure/validator"
	"time"

	"gorm.io/gorm"
)

type RoleService struct {
	db       *gorm.DB
	eventBus *eventbus.EventBus
	validate *validator.Validator
}

func NewRoleService(db *gorm.DB, eventBus *eventbus.EventBus, validate *validator.Validator) *RoleService {
	return &RoleService{
		db:       db,
		eventBus: eventBus,
		validate: validate,
	}
}

func (s *RoleService) Create(req *role.CreateRoleRequest) (*role.RoleResponse, error) {
	// Validate request
	if err := s.validate.Validate(req); err != nil {
		return nil, err
	}

	newRole := &role.Role{
		Role: req.Role,
	}

	if err := s.db.Create(newRole).Error; err != nil {
		return nil, err
	}

	// Publish event
	s.eventBus.Publish(events.Event{
		Type:      events.RoleCreated,
		Timestamp: time.Now(),
		Data: events.RoleCreatedEvent{
			ID:        newRole.ID,
			Role:      newRole.Role,
			CreatedAt: newRole.CreatedAt,
		},
	})

	return &role.RoleResponse{
		ID:        newRole.ID,
		Role:      newRole.Role,
		CreatedAt: newRole.CreatedAt,
		UpdatedAt: newRole.UpdatedAt,
	}, nil
}

func (s *RoleService) GetByID(id uint) (*role.RoleResponse, error) {
	var r role.Role
	if err := s.db.First(&r, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	return &role.RoleResponse{
		ID:        r.ID,
		Role:      r.Role,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}, nil
}

func (s *RoleService) GetAll() ([]role.RoleResponse, error) {
	var roles []role.Role
	if err := s.db.Find(&roles).Error; err != nil {
		return nil, err
	}

	var responses []role.RoleResponse
	for _, r := range roles {
		responses = append(responses, role.RoleResponse{
			ID:        r.ID,
			Role:      r.Role,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}

	return responses, nil
}

func (s *RoleService) Update(id uint, req *role.UpdateRoleRequest) (*role.RoleResponse, error) {
	var r role.Role
	if err := s.db.First(&r, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	if req.Role != "" {
		r.Role = req.Role
	}

	if err := s.db.Save(&r).Error; err != nil {
		return nil, err
	}

	// Publish event
	s.eventBus.Publish(events.Event{
		Type:      events.RoleUpdated,
		Timestamp: time.Now(),
		Data: events.RoleUpdatedEvent{
			ID:        r.ID,
			Role:      r.Role,
			UpdatedAt: r.UpdatedAt,
		},
	})

	return &role.RoleResponse{
		ID:        r.ID,
		Role:      r.Role,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}, nil
}

func (s *RoleService) Delete(id uint) error {
	result := s.db.Delete(&role.Role{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("role not found")
	}

	// Publish event
	s.eventBus.Publish(events.Event{
		Type:      events.RoleDeleted,
		Timestamp: time.Now(),
		Data: events.RoleDeletedEvent{
			ID:        id,
			DeletedAt: time.Now(),
		},
	})

	return nil
}
