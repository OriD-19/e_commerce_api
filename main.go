package main

import (
	"net/http"

	"orid19.com/ecommerce/api/app"
	"orid19.com/ecommerce/api/routes"
)

func main() {
	app := app.NewApp()
	mux := routes.SetRoutes(&app)
	http.ListenAndServe(":3000", mux)
}
