package api

import "orid19.com/ecommerce/api/database"

type ApiHandler struct {
	dbStore database.Store
}

func NewApiHandler(dbStore database.Store) ApiHandler {
	return ApiHandler{dbStore: dbStore}
}
