package connector

import (
	"github.com/karim-w/stdlib/httpclient"
)

func (c *connector_) Req(
	path string,
) httpclient.HTTPRequest {
	req := httpclient.Req(
		c.host + path,
	)

	if c.username != "" || c.password != "" {
		req = req.AddBasicAuth(c.username, c.password)
	}

	return req
}
