package index

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/karim-w/nselastic"
)

func (i *index_[T]) Ensure() error {
	res := i.Connector.Req("/" + i.index).Put()
	if res.IsSuccess() {
		return nil
	}

	code := res.GetStatusCode()
	if code == 409 {
		return nil
	}

	body := res.GetBody()

	var errObj nselastic.ElasticError

	if err := json.Unmarshal(body, &errObj); err != nil {
		return errors.New(
			"failed to create index with status code " + strconv.Itoa(
				code,
			) + " and body " + string(
				body,
			),
		)
	}

	if errObj.Err.Type == "resource_already_exists_exception" {
		return nil
	}

	return &errObj
}
