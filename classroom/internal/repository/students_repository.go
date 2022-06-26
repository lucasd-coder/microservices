package repository

import (
	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"gorm.io/gorm"
)

type StudentsRepository struct {
	Connection *gorm.DB
}

func NewStudentsRepository(connectionDb *gorm.DB) *StudentsRepository {
	return &StudentsRepository{
		Connection: connectionDb,
	}
}

func (db *StudentsRepository) ListAllStudents() []*model.User {
	var student []*model.User
	db.Connection.Find(&student)

	return student
}

func (db *StudentsRepository) GetStudentByAuthUserId(authUserId string) *model.User {
	var student model.User
	db.Connection.Find(&student, "auth_user_id = ?", authUserId)
	return &student
}

func (db *StudentsRepository) GetStudentById(id string) *model.User {
	var student model.User
	db.Connection.Find(&student, "id = ?", id)
	return &student
}

func (db *StudentsRepository) Create(user model.User) (*model.User, error) {
	err := db.Connection.Create(&user).Error
	if err != nil {
		return &model.User{}, err
	}

	return &user, err
}
