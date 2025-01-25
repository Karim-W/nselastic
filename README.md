# Nselastic

Nselastic is my **PERSONAL** tool to interact with ElasticSearch.
ngl the elastic search client in go is a bit of a pain to use, so I made this
to make my life easier casue as everyone knows, unnessary abstraction is the
always the way to go.

## Domains

### Connectors

Connectors are the main way to interact with ElasticSearch. They are the
abstraction that allows you to _connect_ to an ElasticSearch instance and build
the base form of the http requests that will be sent to the server.

#### NewConnector

```go
package main

import (
    "github.com/karim-w/nselastic/impl/connector"
)
func main() {
    conn := connector.New("http://localhost:9200","username","password")
}
```

### Index

Indexes are the main way to interact with the data stored in ElasticSearch.
The Index interface provides the basic CRUD operations that you would expect
from a database with the added bonus of being able to search for data.

#### NewIndex

```go
package main

import (
    "github.com/karim-w/nselastic/impl/connector"
    "github.com/karim-w/nselastic/impl/index"
)

func main() {
    conn := connector.New("http://localhost:9200","username","password")
    idx := index.New(conn,"index_name")
}
```

## License

BSD 3-Clause License

## Author

karim-w

## Contributing

Pull requests are welcome. For major changes, please open an issue first to
discuss what you would like to change.
