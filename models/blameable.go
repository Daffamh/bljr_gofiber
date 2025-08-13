package models

import (
	"time"

	"gorm.io/gorm"
)

type Blameable struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedBy uint           `json:"created_by"`
	UpdatedBy uint           `json:"updated_by"`
	Creator   User           `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Updater   User           `gorm:"foreignKey:UpdatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedBy uint           `json:"deleted_by" gorm:"column:deleted_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
}
