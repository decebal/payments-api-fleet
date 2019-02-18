package main

import (
	"github.com/decebal/payments-api-fleet/api/routes"
	"log"
	"net/http"
)

func main() {
	routes.SetupRoutes()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
