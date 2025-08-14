package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Student struct {
	Id           uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string          `json:"name"`
	Address      string          `json:"address"`
	Email        string          `json:"email" gorm:"unique" validate:"required,email"`
	BirthDate    pgtype.Date     `json:"birth_date" gorm:"type:date"`
	StudentGrade *[]StudentGrade `json:"student_grade"`
	Blameable
}
