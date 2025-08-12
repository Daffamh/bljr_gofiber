package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Student struct {
	Id        uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string      `json:"name"`
	Address   string      `json:"address"`
	BirthDate pgtype.Date `json:"birth_date" gorm:"type:date"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	CreatedBy uint        `json:"created_by"`
	UpdatedBy uint        `json:"updated_by"`
	Creator   User        `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Updater   User        `gorm:"foreignKey:UpdatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
