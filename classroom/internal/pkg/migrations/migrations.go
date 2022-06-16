package migrations

import (
	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Course{}, &model.Enrollment{}, &model.User{})
}
