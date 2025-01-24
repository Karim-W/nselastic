package nselastic

import (
	"github.com/karim-w/stdlib/httpclient"
)

type Connector interface {
	Req(path string) httpclient.HTTPRequest
}
