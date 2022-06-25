//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/lucasd-coder/classroom/internal/pkg/database"
	"github.com/lucasd-coder/classroom/internal/repository"
	"github.com/lucasd-coder/classroom/internal/service"
)

func InitializeCoursesRepository() *repository.CoursesRepository {
	wire.Build(database.GetDatabase, repository.NewCoursesRepository)
	return &repository.CoursesRepository{}
}

func InitializeEnrollemtsRepository() *repository.EnrollmentsRepository {
	wire.Build(database.GetDatabase, repository.NewEnrollmentsRepository)
	return &repository.EnrollmentsRepository{}
}

func InitializeStudentsRepository() *repository.StudentsRepository {
	wire.Build(database.GetDatabase, repository.NewStudentsRepository)
	return &repository.StudentsRepository{}
}

func InitializeCoursesService() *service.CoursesService {
	wire.Build(InitializeCoursesRepository, service.NewCoursesService)
	return &service.CoursesService{}
}

func InitializeEnrollemtsService() *service.EnrollmentsService {
	wire.Build(InitializeEnrollemtsRepository, service.NewEnrollmentsService)
	return &service.EnrollmentsService{}
}

func InitializeStudentsService() *service.StudentsService {
	wire.Build(InitializeStudentsRepository, service.NewStudentsService)
	return &service.StudentsService{}
}
