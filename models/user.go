package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primary_key;autoIncrement"`
	Name         string    `json:"name" gorm:"size:100;not null"`
	Email        string    `json:"email" gorm:"size:100;not null;unique"`
	PasswordHash string    `json:"-" gorm:"size:255;not null"`
	Token        string    `json:"token" gorm:"size:255"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
