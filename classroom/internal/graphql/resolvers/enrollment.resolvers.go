package resolvers

import "github.com/lucasd-coder/classroom/internal/graphql/graph/generated"

func (r *Resolver) Enrollment() generated.EnrollmentResolver { return &enrollmentResolver{r} }

type enrollmentResolver struct{ *Resolver }
