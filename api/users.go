package api

import (
	"encoding/json"
	"net/http"

	"orid19.com/ecommerce/api/types"
)

func (api ApiHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// parse the request body
	var ru types.RequestUser

	// grab the body of the json into the RequestUser struct
	err := json.NewDecoder(r.Body).Decode(&ru)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if the user already exists
	exists, err := api.dbStore.DoesUserExist(ru.Username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "user already exists", http.StatusBadRequest)
		return
	}

	// create a new user

	newUser, err := types.NewUser(ru)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// store the user

	err = api.dbStore.CreateUser(newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return a response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created"))
}

func (api ApiHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {

	var ru types.RequestUser

	err := json.NewDecoder(r.Body).Decode(&ru)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if the user exists
	user, err := api.dbStore.GetUser(ru.Username)

	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	// validate the password
	if !types.ValidatePassword(ru.Password, user.PasswordHash) {
		http.Error(w, "invalid password", http.StatusBadRequest)
		return
	}

	// return a jwt to the client
	accessToken := types.CreateToken(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"access_token": accessToken})
}

func (api ApiHandler) ProtectedRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("protected route"))
}
