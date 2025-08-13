package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Student struct {
	Id           uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string         `json:"name"`
	Address      string         `json:"address"`
	Email        string         `json:"email" gorm:"unique" validate:"required,email"`
	BirthDate    pgtype.Date    `json:"birth_date" gorm:"type:date"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	CreatedBy    uint           `json:"created_by"`
	UpdatedBy    uint           `json:"updated_by"`
	Creator      User           `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Updater      User           `gorm:"foreignKey:UpdatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StudentGrade []StudentGrade `json:"student_grade"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	DeletedBy    *uint          `json:"deleted_by,omitempty"`
}
