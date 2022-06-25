package service

import (
	"fmt"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/interfaces"
	"github.com/lucasd-coder/classroom/internal/repository"
)

var ErrEnrollmentNotFound = fmt.Errorf("enrollment not found")

type EnrollmentsService struct {
	EnrollmentsRepository interfaces.EnrollmentsRepository
}

func NewEnrollmentsService(enrollmentsRepository *repository.EnrollmentsRepository) *EnrollmentsService {
	return &EnrollmentsService{
		EnrollmentsRepository: enrollmentsRepository,
	}
}

func (service *EnrollmentsService) GetByCourseAndStudentId(courseId, studentId string) (*model.Enrollment, error) {
	enrollment := service.EnrollmentsRepository.GetByCourseAndStudentId(courseId, studentId)

	if enrollment.ID == "" {
		return &model.Enrollment{}, ErrEnrollmentNotFound
	}

	return enrollment, nil
}

func (service *EnrollmentsService) ListAllEnrollments() []*model.Enrollment {
	return service.EnrollmentsRepository.ListAllEnrollments()
}

func (service *EnrollmentsService) ListEnrollmentsByStudent(studentId string) []*model.Enrollment {
	return service.EnrollmentsRepository.ListEnrollmentsByStudent(studentId)
}

func (service *EnrollmentsService) FindEnrollmentByID(id string) (*model.Enrollment, error) {
	enrollment := service.EnrollmentsRepository.FindEnrollmentByID(id)

	if enrollment.ID == "" {
		return &model.Enrollment{}, ErrEnrollmentNotFound
	}

	return enrollment, nil
}

func (service *EnrollmentsService) CreateEnrollment(courseId, studentId string) (*model.Enrollment, error) {
	enrollment := model.NewEnrollment(courseId, studentId)

	newEnrollment, err := service.EnrollmentsRepository.Create(*enrollment)
	if err != nil {
		return &model.Enrollment{}, err
	}

	return newEnrollment, nil
}
