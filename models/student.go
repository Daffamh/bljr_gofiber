package models

import (
	"time"
)

type Student struct {
	Id        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date" gorm:"type:date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
