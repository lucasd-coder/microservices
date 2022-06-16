package resolvers

import (
	"context"
	"net/http"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) Enrollments(ctx context.Context) ([]*model.Enrollment, error) {
	gc, claims, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return []*model.Enrollment{}, gqlerror.Errorf(err.Error())
	}
	logger.Log.Info(claims.RegisteredClaims.Subject)

	return []*model.Enrollment{}, nil
}

func (r *entityResolver) FindEnrollmentByID(ctx context.Context, id string) (*model.Enrollment, error) {
	gc, claims, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.Enrollment{}, gqlerror.Errorf(err.Error())
	}
	logger.Log.Info(claims.RegisteredClaims.Subject)
	return &model.Enrollment{
		ID: id,
	}, nil
}
