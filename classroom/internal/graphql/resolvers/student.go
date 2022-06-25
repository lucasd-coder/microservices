package resolvers

import (
	"context"
	"net/http"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) Students(ctx context.Context) ([]*model.User, error) {
	gc, _, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return []*model.User{}, gqlerror.Errorf(err.Error())
	}

	return r.StudentService.ListAllStudents(), nil
}

func (r *entityResolver) FindUserByAuthUserID(ctx context.Context, authUserID string) (*model.User, error) {
	gc, claims, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.User{}, gqlerror.Errorf(err.Error())
	}

	student, err := r.StudentService.GetStudentByAuthUserId(claims.RegisteredClaims.Subject)
	if err != nil {
		logger.Log.Error(err)
		return &model.User{}, gqlerror.Errorf(err.Error())
	}

	return student, nil
}

func (r *userResolver) Enrollments(ctx context.Context, obj *model.User) ([]*model.Enrollment, error) {
	gc, _, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return []*model.Enrollment{}, gqlerror.Errorf(err.Error())
	}

	return r.EnrollmentsService.ListEnrollmentsByStudent(obj.ID), nil
}
