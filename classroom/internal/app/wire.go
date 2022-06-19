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

func InitializeCoursesService() *service.CoursesService {
	wire.Build(InitializeCoursesRepository, service.NewCursesService)
	return &service.CoursesService{}
}
