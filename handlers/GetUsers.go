package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aelpxy/api-http/mocks"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader((http.StatusAccepted))
	json.NewEncoder(w).Encode(mocks.User)
}
