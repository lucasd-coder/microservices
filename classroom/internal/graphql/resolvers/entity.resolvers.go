package resolvers

import "github.com/lucasd-coder/classroom/internal/graphql/graph/generated"

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }