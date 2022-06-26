package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	ID        string `json:"id" gorm:"primaryKey;type:uuid"`
	Student   User   `json:"student_id" binding:"required" gorm:"OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:StudentID"`
	StudentID string
	Course    Course `json:"course_id" binding:"required" gorm:"OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:CourseID"`
	CourseID  string

	CanceledAt sql.NullTime `json:"canceledAt"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
}

func (enrollment *Enrollment) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}

func NewEnrollment(courseId, studentId string) *Enrollment {
	return &Enrollment{
		CourseID:  courseId,
		StudentID: studentId,
	}
}

func (Enrollment) IsEntity() {}
