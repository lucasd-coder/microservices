package resolvers

import (
	"context"

	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/lucasd-coder/classroom/internal/graphql/graph/generated"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/lucasd-coder/classroom/internal/service"
	"github.com/lucasd-coder/classroom/internal/tools"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	CousersService *service.CoursesService
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func CheckContext(ctx context.Context) (*gin.Context, *validator.ValidatedClaims, error) {
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

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
