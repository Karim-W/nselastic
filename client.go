package nselastic

import (
	"context"

	"github.com/karim-w/stdlib/httpclient"
)

type Index[T any] interface {
	// Upsert - add or update a document in the index.
	// - ctx: the context of the request.
	// - id: the id of the document.
	// - doc: the document body to be added or updated.
	Upsert(
		ctx context.Context,
		id string,
		doc T,
	) error
	Fetch(
		ctx context.Context,
		id string,
	) (T, error)
	Delete(
		ctx context.Context,
		id string,
	) error
	Query(
		ctx context.Context,
		query string,
	) ([]T, error)
	Ensure() error
}

type Connector interface {
	Req(path string) httpclient.HTTPRequest
}
