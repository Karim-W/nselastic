package index

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/karim-w/gopts"
)

func (i *index_[T]) Fetch(ctx context.Context, id string) (document gopts.Option[T], err error) {
	res := i.Connector.Req("/" + i.index + "/_doc/" + id).
		Get()

	var result _FetchResult[T]

	if !res.IsSuccess() {
		body := res.GetBody()
		code := res.GetStatusCode()

		if err = json.Unmarshal(body, &result); err != nil {
			err = errors.New(
				"failed to fetch document with status code " + strconv.Itoa(
					code,
				) + " and body " + string(
					body,
				),
			)

			return
		}

		if !result.Found {
			document = gopts.None[T]()
			err = nil
			return
		}
	}

	err = json.Unmarshal(res.GetBody(), &result)
	if err != nil {
		return
	}

	document = gopts.Some(result.Source)

	return
}
