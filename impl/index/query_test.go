package index_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/karim-w/gopts"
	"github.com/karim-w/nselastic"
	"github.com/karim-w/nselastic/impl/connector"
	"github.com/karim-w/nselastic/impl/index"
)

func TestQuery(t *testing.T) {
	type Source struct {
		Id       string `json:"id"`
		Date     string `json:"date"`
		Status   string `json:"status"`
		Type     string `json:"type"`
		Title    string `json:"title"`
		Deleted  bool   `json:"deleted"`
		Archived bool   `json:"archived"`
	}

	HOST, USERNAME, PASSWORD := os.Getenv(
		"ELASTICSEARCH_HOST",
	), os.Getenv(
		"ELASTICSEARCH_USERNAME",
	), os.Getenv(
		"ELASTICSEARCH_PASSWORD",
	)

	if HOST == "" {
		t.Skip("missing environment variables")
	}

	connector := connector.New(
		HOST,
		USERNAME,
		PASSWORD,
	)

	idx, err := index.New[Source](connector, uuid.NewString())
	if err != nil {
		t.Fatal(err)
	}

	d := Source{
		Id:       uuid.NewString(),
		Date:     "2023-01-01",
		Status:   "active",
		Type:     "user",
		Title:    "example",
		Deleted:  false,
		Archived: false,
	}

	// the create part
	if err := idx.Upsert(context.TODO(), d.Id, d); err != nil {
		t.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	// the fetch part
	res, err := idx.Query(context.TODO(), nselastic.Queryable{
		From: gopts.Some(0),
		Size: gopts.Some(1000),
		SortingOptions: []nselastic.SortingOptions{{
			SortField: gopts.Some("date"),
			SortOrder: gopts.Some("asc"),
		}},
		Search: nselastic.QueryBody{
			Searches: []string{"example"},
			Ranges: []nselastic.Range{{
				Key: "date",
				Gte: gopts.Some[any]("2021-01-01"),
				Lte: gopts.Some[any]("2023-12-31"),
			}},
			Equal: []nselastic.Filters{
				{Key: "status", Value: "active"},
				{Key: "type", Value: "user"},
			},
			NotEqual: []nselastic.Filters{
				{Key: "deleted", Value: true},
				{Key: "archived", Value: true},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if res.Total != 1 {
		t.Fatalf("expected 1 got %d", res.Total)
	}

	fmt.Println(res.Data)

	list := res.Data

	if len(list) != 1 {
		t.Fatalf("expected 1 got %d", len(list))
	}

	if list[0] != d {
		t.Fatalf("expected %v got %v", d, list[0])
	}
}
