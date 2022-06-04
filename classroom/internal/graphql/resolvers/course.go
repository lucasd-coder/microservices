package resolvers

import (
	"context"
	"time"

	"github.com/lucasd-coder/classroom/internal/graphql/model"
)

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return []*model.Course{}, nil
}

func (r *mutationResolver) CreateCourse(ctx context.Context, data model.CreateCourseInput) (*model.Course, error) {
	return &model.Course{
		ID: time.Now().String(),
	}, nil
}

func (r *entityResolver) FindCourseByID(ctx context.Context, id string) (*model.Course, error) {
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
	return &model.Course{
		ID: time.Now().String(),
	}, nil
}
