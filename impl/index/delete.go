package index

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/karim-w/nselastic"
)

func (i *index_[T]) Delete(ctx context.Context, ids ...string) error {
	body := strings.Builder{}

	for _, id := range ids {
		body.WriteString(`{"delete":{"_index":"`)
		body.WriteString(i.index)
		body.WriteString(`","_id":"`)
		body.WriteString(id)
		body.WriteString(`"}}`)
		body.WriteRune('\n')
	}

	res := i.Connector.Req("/_bulk").JSON().AddBodyRaw([]byte(body.String())).Post()

	if !res.IsSuccess() {
		code := res.GetStatusCode()
		bdy := res.GetBody()

		var errObj nselastic.Error
		err := json.Unmarshal(bdy, &errObj)
		if err != nil {
			return errors.New(
				"failed to delete document with status code " + strconv.Itoa(
					code,
				) + " and body " + string(
					bdy,
				),
			)
		}

		return &errObj
	}

	var result _DeleteResult
	err := res.SetResult(&result)
	if err != nil {
		return err
	}

	if result.Errors {
		return errors.New("delete operation failed")
	}

	for _, item := range result.Items {
		if item.Delete.Result != "deleted" {
			return errors.New("delete operation failed")
		}

		if item.Delete.Status != 200 {
			return errors.New("delete operation failed")
		}
	}

	return nil
}
