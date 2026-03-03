package readmodel

import (
	"cqrs-blog/internal/domain/role"
	"time"
)

// RoleReadModel is the MongoDB read model for Role
type RoleReadModel struct {
	ID        uint      `bson:"_id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// ToRoleResponse converts a RoleReadModel to RoleResponse
func (r *RoleReadModel) ToRoleResponse() *role.RoleResponse {
	return &role.RoleResponse{
		ID:        r.ID,
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

// ToRoleResponses converts a slice of RoleReadModel to RoleResponse slice
func ToRoleResponses(roles []RoleReadModel) []role.RoleResponse {
	responses := make([]role.RoleResponse, 0, len(roles))
	for _, r := range roles {
		responses = append(responses, *r.ToRoleResponse())
	}
	return responses
}

// NewRoleReadModel creates a RoleReadModel from a domain Role
func NewRoleReadModel(r *role.Role) *RoleReadModel {
	return &RoleReadModel{
		ID:        r.ID,
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}
