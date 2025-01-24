package index_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/karim-w/nselastic/impl/connector"
	"github.com/karim-w/nselastic/impl/index"
)

func TestFetch(t *testing.T) {
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
	fetched, err := idx.Fetch(context.TODO(), d.Id)
	if err != nil {
		t.Fatal(err)
	}

	// document should be fetched
	if fetched.IsNone() {
		t.Fatalf("expected a result got nothing")
	}

	if fetched.Unwrap() != d {
		t.Fatalf("expected %v got %v", d, fetched)
	}
}

func TestFetchNotFound(t *testing.T) {
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

	// the fetch part
	fetched, err := idx.Fetch(context.TODO(), uuid.NewString())
	if err != nil {
		t.Fatal(err)
	}

	if fetched.IsSome() {
		t.Fatalf("expected None got %v", fetched)
	}
}
