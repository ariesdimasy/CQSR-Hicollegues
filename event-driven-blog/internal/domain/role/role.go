package role

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Role      string         `gorm:"size:50;uniqueIndex;not null" json:"role" validate:"required,min=2,max=50"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateRoleRequest struct {
	Role string `json:"role" validate:"required,min=2,max=50"`
}

type UpdateRoleRequest struct {
	Role string `json:"role" validate:"omitempty,min=2,max=50"`
}

type RoleResponse struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
