package resolvers

import (
	"context"
	"net/http"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/lucasd-coder/classroom/internal/tools"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	gc, claims, err := checkContext(ctx)
	if err != nil {
		logger.Log.Error(err)
		gc.AbortWithStatus(http.StatusUnauthorized)
		return []*model.Course{}, gqlerror.Errorf(err.Error())
	}

	logger.Log.Info(claims.RegisteredClaims.Subject)

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

func checkContext(ctx context.Context) (*gin.Context, *validator.ValidatedClaims, error) {
	gc, err := tools.GinContextFromContext(ctx)
	if err != nil {
		return gc, nil, err
	}

	authHeader := gc.GetHeader("Authorization")

	token, err := tools.EnsureValidToken(authHeader, ctx)
	if err != nil {
		return gc, nil, err
	}

	claims, ok := token.(*validator.ValidatedClaims)
	if !ok {
		logger.Log.Warn("unexpected token format")
		return gc, nil, tools.ErrUnAuthorized
	}

	return gc, claims, err
}
