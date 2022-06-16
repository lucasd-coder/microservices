package tools

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
)

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

var (
	ErrUnAuthorized = fmt.Errorf("unauthorized")
	ErrUnexpected   = fmt.Errorf("unexpected error")
	ErrBadRequest   = fmt.Errorf("bad request")
)

func EnsureValidToken(auth string, ctx context.Context) (interface{}, error) {
	if auth == "" {
		return nil, ErrUnAuthorized
	}

	issuerURL, err := url.Parse(os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {		
		logger.Log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {		
		logger.Log.Fatal("Failed to set up the jwt validator")
	}

	token, err := jwtValidator.ValidateToken(ctx, strings.TrimPrefix(auth, "Bearer "))
	if err != nil {
		logger.Log.Error(err)
		return nil, ErrUnAuthorized
	}

	return token, err
}
