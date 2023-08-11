package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/qomaindo-dev/go-bookstore/app/conf"
	"github.com/qomaindo-dev/go-bookstore/app/database"
	"github.com/qomaindo-dev/go-bookstore/graph"
)

const defaultPort = "8080"

func main() {
	config := &database.Config{
		Host:     conf.DBHost,
		Port:     conf.DBPort,
		Username: conf.DBUsername,
		Password: conf.DBPassword,
		DBName:   conf.DBName,
	}
	db, err := database.NewConnection(config)
	if err != nil {
		panic(err)
	}
	database.Migrate(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// repo := repository.NewBookService(db)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.ResolverBook{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
