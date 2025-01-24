package connector

import (
	"github.com/karim-w/nselastic"
)

type connector_ struct {
	host, username, password string
	hook                     nselastic.Hook
}

func New(host, username, password string) nselastic.Connector {
	return &connector_{
		host:     host,
		username: username,
		password: password,
	}
}
