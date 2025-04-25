package index

import "github.com/karim-w/stdlib/httpclient"

func (i *index_[T]) Client(endpoint string) httpclient.HTTPRequest {
	return i.Connector.Req(endpoint).JSON()
}
