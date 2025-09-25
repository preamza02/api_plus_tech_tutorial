package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/preamza02/pokedex/database"
	"github.com/preamza02/pokedex/graph"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()
	// sqlite := database.InitSqlite()
	pgsql := database.InitPostgres()
	resolver := &graph.Resolver{
		DB: pgsql,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
