package index

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/karim-w/nselastic"
)

func (i *index_[T]) List(
	ctx context.Context,
	ids ...string,
) (list []T, err error) {
	m := map[string]interface{}{
		"ids": ids,
	}

	res := i.Connector.Req("/" + i.index + "/_mget").
		JSON().
		AddBody(m).
		Post()

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

	var result _ListResult[T]
	err = res.SetResult(&result)
	if err != nil {
		return
	}

	list = make([]T, 0, len(result.Docs))

	for _, doc := range result.Docs {
		list = append(list, doc.Source)
	}

	return
}
