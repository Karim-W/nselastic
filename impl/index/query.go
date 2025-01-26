package index

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/karim-w/nselastic"
)

func (i *index_[T]) Query(
	ctx context.Context,
	query nselastic.Queryable,
) (QR nselastic.QueryResult[T], err error) {
	m := build_elastic_query(query)

	res := i.Connector.Req("/" + i.index + "/_search").JSON().AddBody(m).Post()

	if !res.IsSuccess() {
		code := res.GetStatusCode()
		body := res.GetBody()

		var errObj nselastic.Error
		err = json.Unmarshal(body, &errObj)
		if err != nil {
			err = errors.New(
				"failed to fetch documents with status code " + strconv.Itoa(
					code,
				) + " and body " + string(
					body,
				),
			)

			return
		}

		err = &errObj
		return
	}

	var result _SearchResult[T]
	err = res.SetResult(&result)
	if err != nil {
		return
	}

	QR.Data = make([]T, 0, len(result.Hits.Hits))

	for _, hit := range result.Hits.Hits {
		QR.Data = append(QR.Data, hit.Source)
	}

	QR.Total = result.Hits.Total.Value

	return
}

type _SearchResult[T any] struct {
	Took     int64   `json:"took"`
	TimedOut bool    `json:"timed_out"`
	Shards   Shards  `json:"_shards"`
	Hits     Hits[T] `json:"hits"`
}

type Hits[T any] struct {
	Total    Total    `json:"total"`
	MaxScore int64    `json:"max_score"`
	Hits     []Hit[T] `json:"hits"`
}

type Hit[T any] struct {
	Index  string  `json:"_index"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source T       `json:"_source"`
}

type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}
