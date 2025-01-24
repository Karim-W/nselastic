package nselastic

type Queryable struct{}

type TraceInfo struct {
	Index, Id string
	Doc       interface{}
	Query     Queryable
}

type ElasticError struct {
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

func (e *ElasticError) Error() string {
	return e.Err.Reason
}

var _ error = &ElasticError{}
