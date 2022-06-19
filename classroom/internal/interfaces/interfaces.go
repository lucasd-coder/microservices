package interfaces

import "github.com/lucasd-coder/classroom/internal/graphql/model"

type (
	CoursesRepository interface {
		ListAllCourses() []*model.Course
		GetCourseById(id string) *model.Course
		GetCourseBySlug(slug string) *model.Course
		Create(curso model.Course) (*model.Course, error)
	}

	CoursesService interface {
		ListAllCourses() []*model.Course
		GetCourseById(id string) (*model.Course, error)
		GetCourseBySlug(slug string) (*model.Course, error)
		CreateCourse(course model.CreateCourseInput) (*model.Course, error)
	}
)
