package index

type _ListResult[T any] struct {
	Docs []_FetchResult[T] `json:"docs"`
}

type _FetchResult[T any] struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int64  `json:"_version"`
	Found   bool   `json:"found"`
	Source  T      `json:"_source"`
}

type _DeleteResult struct {
	Errors bool   `json:"errors"`
	Took   int64  `json:"took"`
	Items  []Item `json:"items"`
}

type Item struct {
	Delete Delete `json:"delete"`
}

type Delete struct {
	Index       string `json:"_index"`
	ID          string `json:"_id"`
	Version     int64  `json:"_version"`
	Result      string `json:"result"`
	Shards      Shards `json:"_shards"`
	SeqNo       int64  `json:"_seq_no"`
	PrimaryTerm int64  `json:"_primary_term"`
	Status      int64  `json:"status"`
}

type Shards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Failed     int64 `json:"failed"`
}
