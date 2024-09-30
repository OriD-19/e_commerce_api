package routes

import (
	"github.com/go-chi/chi/v5"
	chiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"orid19.com/ecommerce/api/app"
	"orid19.com/ecommerce/api/middleware"
)

func SetRoutes(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	// Configure CORS middleware
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
	r.Use(cors.Handler(corsOptions))

	// Add your routes here
	r.Use(chiddleware.Logger)
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuthenticationMiddleware)
		// TODO Implement all the product methods
		// r.Route("/products", func(r chi.Router) {
		// 	r.Get("/", app.ApiHandler.GetProductsHandler)
		// 	r.Post("/", app.ApiHandler.CreateProductHandler)
		// 	r.Get("/{id}", app.ApiHandler.GetProductHandler)
		// })

		r.Get("/protected", app.ApiHandler.ProtectedRouteHandler)
	})

	r.Post("/user/register", app.ApiHandler.RegisterUserHandler)
	r.Post("/user/login", app.ApiHandler.LoginUserHandler)

	return r
}
