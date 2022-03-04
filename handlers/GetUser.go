package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aelpxy/api-http/mocks"
	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	for _, User := range mocks.User {
		if User.Uuid == id {
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(User)

			break
		}
	}
}
