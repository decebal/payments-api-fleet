package main

import (
	"github.com/decebal/payments-api-fleet/api/graphql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://0.0.0.0:"+port},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Handle("/", handler.Playground("Payments Api", "/graphql"))
	router.Handle("/graphql", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}})))

	log.Printf("connect to http://0.0.0.0:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
