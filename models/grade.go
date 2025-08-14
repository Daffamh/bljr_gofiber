package models

type Grade struct {
	Id           uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string         `json:"name"`
	StudentGrade []StudentGrade `json:"student_grade"`
	Blameable
}
