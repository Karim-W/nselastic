package index_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/karim-w/nselastic/impl/connector"
	"github.com/karim-w/nselastic/impl/index"
)

func TestList(t *testing.T) {
	type dummy struct {
		Id string `json:"id"`
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

	idx, err := index.New[dummy](connector, "test")
	if err != nil {
		t.Fatal(err)
	}

	d := dummy{Id: uuid.NewString()}

	// the create part
	if err := idx.Upsert(context.TODO(), d.Id, d); err != nil {
		t.Fatal(err)
	}

	// the fetch part
	list, err := idx.List(context.TODO(), d.Id)
	if err != nil {
		t.Fatal(err)
	}

	if len(list) != 1 {
		t.Fatalf("expected 1 got %d", len(list))
	}

	if list[0] != d {
		t.Fatalf("expected %v got %v", d, list[0])
	}
}

func TestListAll(t *testing.T) {
	type dummy struct {
		Id string `json:"id"`
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

	idx, err := index.New[dummy](connector, "test")
	if err != nil {
		t.Fatal(err)
	}

	ids := make([]string, 0, 10)
	for i := 0; i < 10; i++ {

		d := dummy{Id: uuid.NewString()}
		ids = append(ids, d.Id)

		// the create part
		if err := idx.Upsert(context.TODO(), d.Id, d); err != nil {
			t.Fatal(err)
		}

	}

	// the fetch part
	list, err := idx.List(context.TODO(), ids...)
	if err != nil {
		t.Fatal(err)
	}

	if len(list) == 0 {
		t.Fatalf("expected a result got nothing")
	}

	if len(list) != 10 {
		t.Fatalf("expected 10 got %d", len(list))
	}

	for _, d := range list {
		found := false
		for _, id := range ids {
			if d.Id == id {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v to be in %v", d, ids)
		}

	}
}
