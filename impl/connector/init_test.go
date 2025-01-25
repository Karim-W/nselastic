package connector_test

import (
	"testing"

	"github.com/karim-w/cafe"
	"github.com/karim-w/nselastic/impl/connector"
)

func TestInit(t *testing.T) {
	c, err := cafe.New(
		cafe.Schema{
			connector.CAFE_CONFIG_KEY: connector.CafeConfig,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	sub, err := c.GetSubSchema(connector.CAFE_CONFIG_KEY)
	if err != nil {
		t.Fatal(err)
	}

	_, err = connector.Init(sub)
	if err != nil {
		t.Fatal(err)
	}
}
