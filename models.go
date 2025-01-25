package nselastic

import "github.com/karim-w/gopts"

// Range is a range filter to be applied on the data.
type Range struct {
	// Key is the field to be ranged.
	Key string
	// Gte is the greater than or equal to value.
	// Lte is the less than or equal to value.
	Gte, Lte gopts.Option[any]
}

// Filters is a filter to be applied on the data.
type Filters struct {
	// Key is the field to be filtered.
	Key string
	// Value is the value to be filtered.
	Value any
}

// QueryBody is the query body to be applied on the data.
type QueryBody struct {
	// Searches is a list of strings to be partially matched.
	Searches []string
	// Ranges is a range filter to be applied on the data.
	Ranges []Range
	// Equal is a list of filters to be applied on the data.
	Equal []Filters
	// NotEqual is a list of filters to be excluded from the data.
	NotEqual []Filters
}

// SortingOptions is a list of sorting options to be applied on the data.
type SortingOptions struct {
	// SortField is the field to be sorted.
	SortField, SortOrder gopts.Option[string]
}

// Queryable is a description of the query you would like to perform on the index data
type Queryable struct {
	// From is the starting index of the data to be fetched. Default is 0.
	// Size is the number of data to be fetched. Default is 1000.
	From, Size gopts.Option[int]
	// SortingOptions is a list of sorting options to be applied on the data.
	SortingOptions []SortingOptions
	// Search is the query body to be applied on the data.
	Search QueryBody
}

type TraceInfo struct {
	Index, Id string
	Doc       interface{}
	Query     Queryable
}

type Error struct {
	Err    ErrorDetails `json:"error"`
	Status int64        `json:"status"`
}

type ErrorDetails struct {
	RootCause []ErrorDetails `json:"root_cause,omitempty"`
	Type      string         `json:"type"`
	Reason    string         `json:"reason"`
	IndexUUID string         `json:"index_uuid"`
	Index     string         `json:"index"`
}

func (e *Error) Error() string {
	return e.Err.Reason
}

var _ error = &Error{}
