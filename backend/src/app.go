package main

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	apiServer "github.com/decebal/payments-api-fleet/api/graphql"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func main() {
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Handle("/", handler.Playground("Payments Api", "/graphql"))
	router.Handle("/graphql",
		handler.GraphQL(apiServer.NewExecutableSchema(apiServer.NewResolver())),
	)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
