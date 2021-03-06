package resolvers

import (
	"context"
	"net/http"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/lucasd-coder/classroom/internal/tools"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	gc, _, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return []*model.Course{}, gqlerror.Errorf(err.Error())
	}

	return r.CousersService.ListAllCourses(), nil
}

func (r *mutationResolver) CreateCourse(ctx context.Context, data model.CreateCourseInput) (*model.Course, error) {
	gc, _, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}

	course, err := r.CousersService.CreateCourse(data.Title)
	if err != nil {
		logger.Log.Error(err)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}

	return course, nil
}

func (r *entityResolver) FindCourseByID(ctx context.Context, id string) (*model.Course, error) {
	gc, _, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}
	course, err := r.CousersService.GetCourseById(id)
	if err != nil {
		logger.Log.Error(err)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}

	return course, nil
}

func (r *queryResolver) Course(ctx context.Context, id string) (*model.Course, error) {
	gc, claims, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}

	student, err := r.StudentService.GetStudentByAuthUserId(claims.RegisteredClaims.Subject)
	if err != nil {
		logger.Log.Error(err)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}

	_, err = r.EnrollmentsService.GetByCourseAndStudentId(id, student.ID)
	if err != nil {
		logger.Log.Error(err)
		return &model.Course{}, gqlerror.Errorf(tools.ErrUnAuthorized.Error())
	}

	course, err := r.CousersService.GetCourseById(id)
	if err != nil {
		logger.Log.Error(err)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}

	return course, nil
}
