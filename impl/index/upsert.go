package index

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/karim-w/nselastic"
)

func (i *index_[T]) Upsert(ctx context.Context, id string, doc T) error {
	byts, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	bodyBuilder := strings.Builder{}

	bodyBuilder.WriteString(`{"doc":`)
	bodyBuilder.Write(byts)
	bodyBuilder.WriteString(`,"doc_as_upsert":true}`)

	res := i.Connector.Req("/" + i.index + "/_update/" + id).
		JSON().
		AddBodyRaw([]byte(bodyBuilder.String())).
		Post()

	if res.IsSuccess() {
		return nil
	}

	body := res.GetBody()
	code := res.GetStatusCode()

	var errObj nselastic.ElasticError

	if err := json.Unmarshal(body, &errObj); err != nil {
		return errors.New(
			"failed to upsert document with status code " + strconv.Itoa(
				code,
			) + " and body " + string(
				body,
			),
		)
	}

	return &errObj
}
