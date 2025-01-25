package index

import (
	"github.com/karim-w/nselastic"
)

type index_[T any] struct {
	Connector nselastic.Connector
	index     string
}

func New[T any](connector nselastic.Connector, index string) (idx nselastic.Index[T], err error) {
	idx = &index_[T]{
		Connector: connector,
		index:     index,
	}

	err = idx.Ensure()
	return
}
