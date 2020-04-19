package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	resolver "github.com/clwiseman/letsgopokemon/internal"
	"github.com/clwiseman/letsgopokemon/internal/generated"
	"github.com/clwiseman/letsgopokemon/internal/pokedex"
)

const (
	defaultPort = "8080"
)

func main() {
	// To connect to GraphQL Playground
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// Connect to database
	pd := pokedex.NewPokedex()
}
