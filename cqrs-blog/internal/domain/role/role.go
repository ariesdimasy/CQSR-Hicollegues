package role

import (
	"time"

	"gorm.io/gorm"
)

// Entity
type Role struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:50;uniqueIndex;not null" json:"name" validate:"required,min=2,max=50"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Response DTO
type RoleResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToRoleResponse(r *Role) *RoleResponse {
	return &RoleResponse{
		ID:        r.ID,
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func ToRoleResponses(roles []Role) []RoleResponse {
	responses := make([]RoleResponse, 0, len(roles))
	for _, r := range roles {
		responses = append(responses, *ToRoleResponse(&r))
	}
	return responses
}
