package handlers

import (
	"cqrs-blog/internal/cqrs/queries"
	"cqrs-blog/internal/domain/role"
	"errors"

	"gorm.io/gorm"
)

// RoleQueryHandler handles all role-related queries
type RoleQueryHandler struct {
	db *gorm.DB
}

// NewRoleQueryHandler creates a new RoleQueryHandler
func NewRoleQueryHandler(db *gorm.DB) *RoleQueryHandler {
	return &RoleQueryHandler{db: db}
}

// HandleGetByID handles the GetRoleByIDQuery
func (h *RoleQueryHandler) HandleGetByID(query queries.GetRoleByIDQuery) (*role.RoleResponse, error) {
	var r role.Role
	if err := h.db.First(&r, query.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	return role.ToRoleResponse(&r), nil
}

// HandleGetAll handles the GetAllRolesQuery
func (h *RoleQueryHandler) HandleGetAll(query queries.GetAllRolesQuery) ([]role.RoleResponse, error) {
	var roles []role.Role
	if err := h.db.Find(&roles).Error; err != nil {
		return nil, err
	}

	return role.ToRoleResponses(roles), nil
}
