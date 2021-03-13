# gqltest [![Go Reference](https://pkg.go.dev/badge/github.com/tsh96/gqltest.svg)](https://pkg.go.dev/github.com/tsh96/gqltest)

gqltest is a simple test library inspired by httptest. It provide a thin abstraction layer from httptest.

## Installation

Use go get to install gqltest

```bash
go get github.com/tsh96/gqltest
```

## Usage

```go
package test

import (
  "encoding/json"
  "fmt"
  "testing"

  "github.com/99designs/gqlgen/graphql/handler"
  "github.com/tsh96/gqltest"
  "github.com/tsh96/gqltest/example/graph"
  "github.com/tsh96/gqltest/example/graph/generated"
  "github.com/tsh96/gqltest/example/graph/model"
)

func TestQuery(t *testing.T) {
  // This example use gqlgen as a graphql backend
  srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

  // Create a new recorder
  w := gqltest.NewRecorder()

  // Create a new request from file
  r, _ := gqltest.NewRequestFromFile("query.gql")
  // Instead of create from file, it also can be created inline
  // r, _ := gqltest.NewRequest(`
  //   query getTodo{
  //     todos{
  //       id
  //     }
  //   }`)

  // run server
  srv.ServeHTTP(w, r)

  // Get the response body
  res, err := w.ResponseBody()

  // Error handling
  if err != nil {
    t.Error(err)
  }
  if res.Errors != nil {
    t.Error(res.Errors)
  }

  // Un-marshal data
  data := map[string]model.Todo{}
  json.Unmarshal(res.Data, &data)

  // do other things with data
  fmt.Println(data["getTodo"])
}
```

## Features

- [x] Query
- [x] Mutation
- [x] Subscription
- [x] Variables (Available in gqltest.RequestOption)
- [x] Headers (Inherit from httptest.Request)
- [ ] Files Upload

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
