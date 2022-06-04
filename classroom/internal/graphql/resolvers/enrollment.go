package resolvers

import (
	"context"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
)

func (r *queryResolver) Enrollments(ctx context.Context) ([]*model.Enrollment, error) {
	return []*model.Enrollment{}, nil
}

func (r *entityResolver) FindEnrollmentByID(ctx context.Context, id string) (*model.Enrollment, error) {
	return &model.Enrollment{
		ID: id,
	}, nil
}
