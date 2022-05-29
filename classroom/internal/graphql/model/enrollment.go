package model

import "time"

type Enrollment struct {
	ID        string `json:"id"`
	Student   User   `json:"student_id"`
	StudentID string
	Course    Course `json:"course_id"`
	CouseID   string

	CanceleAt time.Time `json:"canceleAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (Enrollment) IsEntity() {}
