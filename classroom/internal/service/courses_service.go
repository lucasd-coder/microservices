package service

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/interfaces"
	"github.com/lucasd-coder/classroom/internal/repository"
)

type CoursesService struct {
	CoursesRepository interfaces.CoursesRepository
}

var (
	ErrCourseNotFound     = fmt.Errorf("course not found")
	ErrCourseAlreadyExist = fmt.Errorf("course already exists")
)

func NewCoursesService(coursesRepository *repository.CoursesRepository) *CoursesService {
	return &CoursesService{
		CoursesRepository: coursesRepository,
	}
}

func (service *CoursesService) ListAllCourses() []*model.Course {
	return service.CoursesRepository.ListAllCourses()
}

func (service *CoursesService) GetCourseById(id string) (*model.Course, error) {
	course := service.CoursesRepository.GetCourseById(id)

	if course.ID == "" {
		return &model.Course{}, ErrCourseNotFound
	}

	return course, nil
}

func (service *CoursesService) GetCourseBySlug(slug string) (*model.Course, error) {
	course := service.CoursesRepository.GetCourseBySlug(slug)

	if course.ID == "" {
		return &model.Course{}, ErrCourseNotFound
	}

	return course, nil
}

func (service *CoursesService) CreateCourse(title string) (*model.Course, error) {
	slug := slug.Make(title)

	aux := service.CoursesRepository.GetCourseBySlug(slug)

	if aux.ID != "" {
		return &model.Course{}, ErrCourseAlreadyExist
	}

	course := model.NewCourse(title, slug)

	newCourse, err := service.CoursesRepository.Create(*course)
	if err != nil {
		return &model.Course{}, err
	}

	return newCourse, nil
}
