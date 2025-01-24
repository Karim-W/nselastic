package nselastic

import (
	"context"

	"github.com/karim-w/gopts"
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
	// Fetch - fetches a document from the index.
	// - ctx: the context of the request.
	// - id: the id of the document.
	// Returns:
	// - document: an option of the document, can be None if not found.
	// - error: an error if any. Not found is not considered an error and is represented by the None option
	Fetch(
		ctx context.Context,
		id string,
	) (gopts.Option[T], error)
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
