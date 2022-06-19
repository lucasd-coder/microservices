package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (course *Course) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}

func NewCourse(title, slug string) *Course {
	return &Course{
		Title: title,
		Slug:  slug,
	}
}

func (Course) IsEntity() {}

type CreateCourseInput struct {
	Title string `json:"title"`
}
