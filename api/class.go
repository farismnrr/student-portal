/**
 * This function handles the endpoint to fetch all classes.
 * It calls the FetchAll method from the classService to retrieve all classes.
 * If an error occurs during the fetch operation, it returns a 500 Internal Server Error response.
 * Otherwise, it returns a 200 OK response with the list of classes encoded in JSON format.
 */

package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) FetchAllClass(w http.ResponseWriter, r *http.Request) {
	classes, err := api.classService.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(classes)
}
