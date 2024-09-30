package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"orid19.com/ecommerce/api/types"
)

func (api ApiHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	// get the products from the database

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		http.Error(w, "offset query parameter is required", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		http.Error(w, "limit query parameter is required", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pr := types.ProductsRequest{
		Offset: offset,
		Limit:  limit,
	}

	val := validator.New()
	val.Struct(pr)

	products, err := api.dbStore.GetProducts(pr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the products
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (api ApiHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {

	queryId := r.URL.Query().Get("id")

	if queryId == "" {
		http.Error(w, "specify an id for the product", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(queryId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := api.dbStore.GetProduct(id)

	if err != nil {
		http.Error(w, fmt.Sprintf("product not found: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (api ApiHandler) InsertProductHandler(w http.ResponseWriter, r *http.Request) {
	var product types.InsertProductRequest

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = api.dbStore.InsertProduct(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("product created"))
}
