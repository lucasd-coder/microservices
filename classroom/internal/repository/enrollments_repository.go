package repository

import (
	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"gorm.io/gorm"
)

type EnrollmentsRepository struct {
	Connection *gorm.DB
}

func NewEnrollmentsRepository(connectionDb *gorm.DB) *EnrollmentsRepository {
	return &EnrollmentsRepository{
		Connection: connectionDb,
	}
}

func (db *EnrollmentsRepository) FindEnrollmentByID(id string) *model.Enrollment {
	var enrollment model.Enrollment
	db.Connection.Find(&enrollment, "id = ?", id)
	return &enrollment
}

func (db *EnrollmentsRepository) GetByCourseAndStudentId(courseId, studentId string) *model.Enrollment {
	var enrollment model.Enrollment
	db.Connection.Where("course_id = ? AND student_id = ? AND canceled_at IS NULL", courseId, studentId).Find(&enrollment)
	return &enrollment
}

func (db *EnrollmentsRepository) ListAllEnrollments() []*model.Enrollment {
	var enrollment []*model.Enrollment
	db.Connection.Where("canceled_at IS NULL").Order("created_at desc").Find(&enrollment)

	return enrollment
}

func (db *EnrollmentsRepository) ListEnrollmentsByStudent(studentId string) []*model.Enrollment {
	var enrollment []*model.Enrollment
	db.Connection.Preload("Student").Where("student_id = ? AND canceled_at IS NULL", studentId).Find(&enrollment)
	return enrollment
}

func (db *EnrollmentsRepository) Create(enrollment model.Enrollment) (*model.Enrollment, error) {
	err := db.Connection.Create(&enrollment).Error
	if err != nil {
		return &model.Enrollment{}, err
	}
	return &enrollment, err
}
