package models

type StudentGrade struct {
	Id         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	StudentId  uint      `json:"student_id" gorm:"not null"`
	GradeId    uint      `json:"grade_id"  gorm:"not null"`
	HomeRoomId uint      `json:"home_room_id"  gorm:"not null"`
	StatusId   uint      `json:"status_id"  gorm:"not null"`
	Student    *Student  `json:"student" gorm:"foreignKey:StudentId;constraint:OnUpdate:CASCADE;"`
	Grade      *Grade    `json:"grade" gorm:"foreignKey:GradeId;constraint:OnUpdate:CASCADE;"`
	HomeRoom   *HomeRoom `json:"home_room" gorm:"foreignKey:HomeRoomId;constraint:OnUpdate:CASCADE;"`
	Status     *Status   `json:"status" gorm:"foreignKey:StatusId;constraint:OnUpdate:CASCADE;"`
	Blameable
}
