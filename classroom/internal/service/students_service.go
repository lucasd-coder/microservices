package service

import (
	"fmt"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/interfaces"
	"github.com/lucasd-coder/classroom/internal/repository"
)

var ErrStudentNotFound = fmt.Errorf("student not found")

type StudentsService struct {
	StudentsRepository interfaces.StudentsRepository
}

func NewStudentsService(studentsRepository *repository.StudentsRepository) *StudentsService {
	return &StudentsService{
		StudentsRepository: studentsRepository,
	}
}

func (service *StudentsService) ListAllStudents() []*model.User {
	return service.StudentsRepository.ListAllStudents()
}

func (service *StudentsService) GetStudentByAuthUserId(authUserId string) (*model.User, error) {
	student := service.StudentsRepository.GetStudentByAuthUserId(authUserId)

	if student.ID == "" {
		return &model.User{}, ErrStudentNotFound
	}

	return student, nil
}

func (service *StudentsService) GetStudentById(id string) (*model.User, error) {
	student := service.StudentsRepository.GetStudentById(id)

	if student.ID == "" {
		return &model.User{}, ErrStudentNotFound
	}

	return student, nil
}

func (service *StudentsService) CreateStudent(authUserId string) (*model.User, error) {
	student := model.NewUser(authUserId)

	newStudent, err := service.StudentsRepository.Create(*student)
	if err != nil {
		return &model.User{}, err
	}

	return newStudent, nil
}
