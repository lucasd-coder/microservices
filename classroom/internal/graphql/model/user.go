package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID string `json:"id" gorm:"primaryKey;type:uuid"`

	AuthUserId string `json:"authUserId;"gorm:"unique"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Enrollments []*Enrollment `json:"enrollments" gorm:"foreignKey:StudentID"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}

func (User) IsEntity() {}
