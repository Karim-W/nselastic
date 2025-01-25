package index_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/karim-w/nselastic/impl/connector"
	"github.com/karim-w/nselastic/impl/index"
)

func TestDeleteOne(t *testing.T) {
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

	err = idx.Delete(context.TODO(), d.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteMany(t *testing.T) {
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

	err = idx.Delete(context.TODO(), ids...)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteManyWithADummy(t *testing.T) {
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

	ids := make([]string, 0, 11)
	for i := 0; i < 10; i++ {
		d := dummy{Id: uuid.NewString()}

		ids = append(ids, d.Id)

		// the create part
		if err := idx.Upsert(context.TODO(), d.Id, d); err != nil {
			t.Fatal(err)
		}
	}

	ids = append(ids, "dummy")

	err = idx.Delete(context.TODO(), ids...)
	if err == nil {
		t.Fatal("expected an error")
	}
}
