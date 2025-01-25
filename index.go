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
	// Delete - deletes one or many documents from the index.
	// - ctx: the context of the request.
	// - ids: the ids of the documents to be deleted. multi argument field for multiple ids.
	Delete(
		ctx context.Context,
		ids ...string,
	) error
	// List - retrieves one or many documents from the index.
	// - ctx: the context of the request.
	// - ids: the ids of the documents to be retrieved. multi argument field for multiple ids.
	// Returns:
	// - list: a list of documents.
	// - error: an error if any.
	List(
		ctx context.Context,
		ids ...string,
	) ([]T, error)
	Query(
		ctx context.Context,
		query string,
	) ([]T, error)
	// Ensure - ensures the index exists.
	// Returns:
	// - error: an error if any.
	Ensure() error
}
