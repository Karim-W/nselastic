package connector

import (
	"github.com/karim-w/cafe"
	"github.com/karim-w/nselastic"
)

const CAFE_CONFIG_KEY = "nselastic"

var CafeConfig = cafe.SubSchema(
	CAFE_CONFIG_KEY,
	cafe.Schema{
		"ELASTICSEARCH_HOST":     cafe.String("ELASTICSEARCH_HOST").Require(),
		"ELASTICSEARCH_USERNAME": cafe.String("ELASTICSEARCH_USERNAME"),
		"ELASTICSEARCH_PASSWORD": cafe.String("ELASTICSEARCH_PASSWORD"),
	},
)

func Init(
	subschema *cafe.Cafe,
) (conn nselastic.Connector, err error) {
	var (
		host     string
		username string
		password string
	)

	host, err = subschema.GetString("ELASTICSEARCH_HOST")
	if err != nil {
		return
	}

	username, err = subschema.GetString("ELASTICSEARCH_USERNAME")
	if err != nil {
		return
	}

	password, err = subschema.GetString("ELASTICSEARCH_PASSWORD")
	if err != nil {
		return
	}

	conn = New(
		host,
		username,
		password,
	)

	return
}
