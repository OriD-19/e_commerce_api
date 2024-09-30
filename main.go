package main

import (
	"fmt"
	"net/http"
	"os"

	"orid19.com/ecommerce/api/app"
	"orid19.com/ecommerce/api/routes"
)

func main() {
	app := app.NewApp()
	mux := routes.SetRoutes(&app)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
