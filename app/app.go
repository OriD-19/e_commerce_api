package app

import (
	"orid19.com/ecommerce/api/api"
	"orid19.com/ecommerce/api/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	db := database.NewDB()
	apiHandler := api.NewApiHandler(db)

	return App{
		ApiHandler: apiHandler,
	}
}
