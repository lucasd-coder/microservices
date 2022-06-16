package resolvers

import (
	"context"
	"net/http"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) Students(ctx context.Context) ([]*model.User, error) {
	gc, claims, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return []*model.User{}, gqlerror.Errorf(err.Error())
	}
	logger.Log.Info(claims.RegisteredClaims.Subject)
	return []*model.User{}, nil
}

func (r *entityResolver) FindUserByAuthUserID(ctx context.Context, authUserID string) (*model.User, error) {
	gc, claims, err := CheckContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return &model.User{}, gqlerror.Errorf(err.Error())
	}
	logger.Log.Info(claims.RegisteredClaims.Subject)
	return &model.User{
		AuthUserId: authUserID,
	}, nil
}
