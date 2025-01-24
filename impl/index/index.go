package index

import (
	"context"

	"github.com/karim-w/nselastic"
)

type index_[T any] struct {
	Connector nselastic.Connector
	index     string
}

// Delete implements nselastic.Index.
func (i *index_[T]) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Query implements nselastic.Index.
func (i *index_[T]) Query(ctx context.Context, query string) ([]T, error) {
	panic("unimplemented")
}

func New[T any](connector nselastic.Connector, index string) (idx nselastic.Index[T], err error) {
	idx = &index_[T]{
		Connector: connector,
		index:     index,
	}

	err = idx.Ensure()
	return
}
