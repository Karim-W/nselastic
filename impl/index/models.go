package index

type _FetchResult[T any] struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int64  `json:"_version"`
	Found   bool   `json:"found"`
	Source  T      `json:"_source"`
}
