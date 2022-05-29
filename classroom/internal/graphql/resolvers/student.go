package resolvers

import (
	"context"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
)

func (r *queryResolver) Students(ctx context.Context) ([]*model.User, error) {
	return []*model.User{}, nil
}

func (r *entityResolver) FindUserByAuthUserID(ctx context.Context, authUserID string) (*model.User, error) {
	return &model.User{
		AuthUserId: authUserID,
	}, nil
}
