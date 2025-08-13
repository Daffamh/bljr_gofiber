package models

import (
	"gorm.io/gorm"
	"time"
)

type StudentGrade struct {
	Id         uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	StudentId  uint           `json:"student_id" gorm:"not null"`
	GradeId    uint           `json:"grade_id"  gorm:"not null"`
	HomeRoomId uint           `json:"home_room_id"  gorm:"not null"`
	StatusId   uint           `json:"status_id"  gorm:"not null"`
	CreatedBy  uint           `json:"created_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedBy  uint           `json:"updated_by"`
	UpdatedAt  time.Time      `json:"updated_at"`
	Creator    User           `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Updater    User           `gorm:"foreignKey:UpdatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Student    Student        `json:"student" gorm:"foreignKey:StudentId;constraint:OnUpdate:CASCADE;"`
	Grade      Grade          `json:"grade" gorm:"foreignKey:GradeId;constraint:OnUpdate:CASCADE;"`
	HomeRoom   HomeRoom       `json:"home_room" gorm:"foreignKey:HomeRoomId;constraint:OnUpdate:CASCADE;"`
	Status     Status         `json:"status" gorm:"foreignKey:StatusId;constraint:OnUpdate:CASCADE;"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	DeletedBy  *uint          `json:"deleted_by,omitempty"`
}
