package index

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/karim-w/gopts"
	"github.com/karim-w/nselastic"
)

func TestQueryBuilder(t *testing.T) {
	m := build_elastic_query(nselastic.Queryable{
		From: gopts.Some(0),
		Size: gopts.Some(1000),
		Search: nselastic.QueryBody{
			Searches: []string{"example"},
			Ranges: []nselastic.Range{
				{
					Key: "date",
					Gte: gopts.Some[any]("2023-01-01"),
					Lte: gopts.Some[any]("2023-12-31"),
				},
			},
			Equal: []nselastic.Filters{
				{
					Key:   "status",
					Value: "active",
				},
				{
					Key:   "type",
					Value: "user",
				},
			},
			NotEqual: []nselastic.Filters{
				{
					Key:   "deleted",
					Value: true,
				},
				{
					Key:   "archived",
					Value: true,
				},
			},
		},
	})

	bytes, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	fmt.Println(string(bytes))
}
