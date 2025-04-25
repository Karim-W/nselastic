package index_test

import (
	"os"
	"testing"

	"github.com/karim-w/nselastic/impl/connector"
	"github.com/karim-w/nselastic/impl/index"
)

func TestGetClient(t *testing.T) {
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

	idx, err := index.New[any](connector, "test")
	if err != nil {
		t.Fatal(err)
	}

	res := idx.Client("/_cat/indices?v").Get()
	if !res.IsSuccess() {
		t.Error("failed to call api")
	}
}
