package models

import (
	"gorm.io/gorm"
	"time"
)

type HomeRoom struct {
	Id           uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string         `json:"name"`
	CreatedBy    uint           `json:"created_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedBy    uint           `json:"updated_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Creator      User           `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Updater      User           `gorm:"foreignKey:UpdatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StudentGrade []StudentGrade `json:"student_grade"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	DeletedBy    *uint          `json:"deleted_by,omitempty"`
}
