package resolvers

import (
	"context"
	"net/http"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) Enrollments(ctx context.Context) ([]*model.Enrollment, error) {
	gc, _, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return []*model.Enrollment{}, gqlerror.Errorf(err.Error())
	}

	return r.EnrollmentsService.ListAllEnrollments(), nil
}

func (r *entityResolver) FindEnrollmentByID(ctx context.Context, id string) (*model.Enrollment, error) {
	gc, _, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.Enrollment{}, gqlerror.Errorf(err.Error())
	}

	enrollment, err := r.EnrollmentsService.FindEnrollmentByID(id)
	if err != nil {
		logger.Log.Error(err)
		return &model.Enrollment{}, gqlerror.Errorf(err.Error())
	}

	return enrollment, nil
}

func (r *enrollmentResolver) Student(ctx context.Context, obj *model.Enrollment) (*model.User, error) {
	gc, _, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.User{}, gqlerror.Errorf(err.Error())
	}

	student, err := r.StudentService.GetStudentById(obj.StudentID)
	if err != nil {
		logger.Log.Error(err)
		return &model.User{}, gqlerror.Errorf(err.Error())
	}

	return student, nil
}

func (r *enrollmentResolver) Course(ctx context.Context, obj *model.Enrollment) (*model.Course, error) {
	gc, _, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}
	course, err := r.CousersService.GetCourseById(obj.CourseID)
	if err != nil {
		logger.Log.Error(err)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}

	return course, nil
}
