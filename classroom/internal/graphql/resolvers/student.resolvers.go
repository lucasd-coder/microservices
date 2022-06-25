package resolvers

import "github.com/lucasd-coder/classroom/internal/graphql/graph/generated"

func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
