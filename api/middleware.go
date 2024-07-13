
/**
* This file contains middleware functions for handling authentication and HTTP method checks.
* The Auth middleware function checks for a session token in the request cookie, validates the token, and adds the username to the request context.
* The Get middleware function checks if the HTTP method is GET, otherwise returns a method not allowed error.
* The Post middleware function checks if the HTTP method is POST, otherwise returns a method not allowed error.
* The Delete middleware function checks if the HTTP method is DELETE, otherwise returns a method not allowed error.
* The Put middleware function checks if the HTTP method is PUT, otherwise returns a method not allowed error.
*
* Middleware Functions:
* - Auth: Handles authentication by checking session token and adding username to context.
* - Get: Handles GET method check.
* - Post: Handles POST method check.
* - Delete: Handles DELETE method check.
* - Put: Handles PUT method check.
*
* Routes/Endpoints:
* - No specific routes/endpoints defined in this file.
*/

package api

import (
	"a21hc3NpZ25tZW50/model"
	"context"
	"encoding/json"
	"net/http"
)

func (api *API) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}
		session := c.Value

		sessionFound, err := api.sessionService.TokenValidity(session)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		type contextKey string
		ctx := context.WithValue(r.Context(), contextKey("username"), sessionFound.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Put(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
