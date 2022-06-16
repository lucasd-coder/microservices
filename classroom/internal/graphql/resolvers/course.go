package resolvers

import (
	"context"
	"net/http"
	"time"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	gc, claims, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return []*model.Course{}, gqlerror.Errorf(err.Error())
	}

	logger.Log.Info(claims.RegisteredClaims.Subject)

	return []*model.Course{}, nil
}

func (r *mutationResolver) CreateCourse(ctx context.Context, data model.CreateCourseInput) (*model.Course, error) {
	gc, claims, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}

	logger.Log.Info(claims.RegisteredClaims.Subject)
	return &model.Course{
		ID: time.Now().String(),
	}, nil
}

func (r *entityResolver) FindCourseByID(ctx context.Context, id string) (*model.Course, error) {
	gc, claims, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}
	logger.Log.Info(claims.RegisteredClaims.Subject)

	title := "Course " + id
	if id == "1234" {
		title = "Me"
	}

	return &model.Course{
		ID:    id,
		Title: title,
	}, nil
}

func (r *queryResolver) Course(ctx context.Context, id string) (*model.Course, error) {
	gc, claims, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.Course{}, gqlerror.Errorf(err.Error())
	}
	logger.Log.Info(claims.RegisteredClaims.Subject)

	return &model.Course{
		ID: time.Now().String(),
	}, nil
}
