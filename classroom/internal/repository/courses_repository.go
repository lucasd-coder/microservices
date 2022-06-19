package repository

import (
	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"gorm.io/gorm"
)

type CoursesRepository struct {
	Connection *gorm.DB
}

func NewCoursesRepository(connectionDb *gorm.DB) *CoursesRepository {
	return &CoursesRepository{
		Connection: connectionDb,
	}
}

func (db *CoursesRepository) ListAllCourses() []*model.Course {
	var cursos []*model.Course
	db.Connection.Find(&cursos)

	return cursos
}

func (db *CoursesRepository) GetCourseById(id string) *model.Course {
	var curso model.Course
	db.Connection.Find(&curso, "id = ?", id)

	return &curso
}

func (db *CoursesRepository) GetCourseBySlug(slug string) *model.Course {
	var curso model.Course
	db.Connection.Find(&curso, "slug = ?", slug)

	return &curso
}

func (db *CoursesRepository) Create(course model.Course) (*model.Course, error) {
	err := db.Connection.Create(&course).Error
	if err != nil {
		return &model.Course{}, err
	}
	return &course, err
}
