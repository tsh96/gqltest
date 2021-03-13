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
	r, _ := gqltest.NewRequestFromFile("query.gql", nil)
	// Instead of create from file, it also can be created inline
	// r, _ := gqltest.NewRequest(`
	// 	query getTodo{
	// 		todos{
	// 			id
	// 		}
	// 	}`)

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

func TestMutation(t *testing.T) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	w := gqltest.NewRecorder()
	r, _ := gqltest.NewRequestFromFile("mutation.gql", nil)

	srv.ServeHTTP(w, r)
	rb, _ := w.ResponseBody()
	a := map[string]model.Todo{}
	json.Unmarshal(rb.Data, &a)
	fmt.Println(a["createTodo"])
}
