package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint           `json:"id" gorm:"primary_key;autoIncrement"`
	Name         string         `json:"name" gorm:"size:100;not null"`
	Email        string         `json:"email" gorm:"size:100;not null;unique"`
	PasswordHash string         `json:"-" gorm:"size:255;not null"`
	Token        string         `json:"-" gorm:"size:255"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	DeletedBy    *uint          `json:"deleted_by,omitempty"`
}
