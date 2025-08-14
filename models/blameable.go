package models

import (
	"time"

	"gorm.io/gorm"
)

type Blameable struct {
	CreatedById uint      `json:"created_by_id"`
	CreatedBy   User      `json:"created_by" gorm:"foreignKey:CreatedById;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time `json:"created_at"`

	UpdatedById uint      `json:"updated_by_id"`
	UpdatedBy   User      `json:"updated_by" gorm:"foreignKey:UpdatedById;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedAt   time.Time `json:"updated_at"`

	DeletedById uint           `json:"deleted_by_id" gorm:"column:deleted_by_id"`
	DeletedBy   User           `json:"deleted_by" gorm:"foreignKey:DeletedById;"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
}
