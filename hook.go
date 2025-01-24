package nselastic

import (
	"context"
	"time"
)

type Hook interface {
	Before(
		ctx context.Context,
		index string,
		id string,
		doc interface{},
		query Queryable,
	) error
	After(
		ctx context.Context,
		index string,
		id string,
		doc interface{},
		query Queryable,
		start time.Time,
		end time.Time,
		status int,
		err error,
	) error
}
