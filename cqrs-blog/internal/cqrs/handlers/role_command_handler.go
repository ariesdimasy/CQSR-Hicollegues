package handlers

import (
	"cqrs-blog/internal/cqrs/commands"
	"cqrs-blog/internal/domain/role"
	"cqrs-blog/internal/infrastructure/sync"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// RoleCommandHandler handles all role-related commands (Write → PostgreSQL)
type RoleCommandHandler struct {
	db          *gorm.DB
	validate    *validator.Validate
	syncService *sync.SyncService
}

// NewRoleCommandHandler creates a new RoleCommandHandler
func NewRoleCommandHandler(db *gorm.DB, validate *validator.Validate, syncService *sync.SyncService) *RoleCommandHandler {
	return &RoleCommandHandler{
		db:          db,
		validate:    validate,
		syncService: syncService,
	}
}

// HandleCreate handles the CreateRoleCommand
func (h *RoleCommandHandler) HandleCreate(cmd commands.CreateRoleCommand) (*role.RoleResponse, error) {
	if err := h.validate.Struct(cmd); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	newRole := &role.Role{
		Name: cmd.Name,
	}

	if err := h.db.Create(newRole).Error; err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}

	// Sync to MongoDB (read side)
	if err := h.syncService.SyncRole(newRole); err != nil {
		// Log but don't fail the request — write succeeded
		fmt.Printf("WARNING: failed to sync role to read DB: %v\n", err)
	}

	return role.ToRoleResponse(newRole), nil
}

// HandleUpdate handles the UpdateRoleCommand
func (h *RoleCommandHandler) HandleUpdate(cmd commands.UpdateRoleCommand) (*role.RoleResponse, error) {
	if err := h.validate.Struct(cmd); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	var r role.Role
	if err := h.db.First(&r, cmd.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	if cmd.Name != "" {
		r.Name = cmd.Name
	}

	if err := h.db.Save(&r).Error; err != nil {
		return nil, fmt.Errorf("failed to update role: %w", err)
	}

	// Sync to MongoDB (read side)
	if err := h.syncService.SyncRole(&r); err != nil {
		fmt.Printf("WARNING: failed to sync role to read DB: %v\n", err)
	}

	return role.ToRoleResponse(&r), nil
}

// HandleDelete handles the DeleteRoleCommand
func (h *RoleCommandHandler) HandleDelete(cmd commands.DeleteRoleCommand) error {
	result := h.db.Delete(&role.Role{}, cmd.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("role not found")
	}

	// Sync to MongoDB (read side)
	if err := h.syncService.DeleteRole(cmd.ID); err != nil {
		fmt.Printf("WARNING: failed to delete role from read DB: %v\n", err)
	}

	return nil
}
