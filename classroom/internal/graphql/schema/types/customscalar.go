package types

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalNullTime(ns sql.NullTime) graphql.Marshaler {
	if !ns.Valid {
		// this is also important, so we can detect if this scalar is used in a not null context and return an appropriate error
		return graphql.Null
	}
	return graphql.MarshalTime(ns.Time)
}

func UnmarshalNullTime(v interface{}) (sql.NullTime, error) {
	if v == nil {
		return sql.NullTime{Valid: false}, nil
	}
	// again you can delegate to the default implementation to save yourself some work.
	s, err := graphql.UnmarshalTime(v)
	return sql.NullTime{Time: s, Valid: true}, err
}
