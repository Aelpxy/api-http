package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aelpxy/api-http/mocks"
	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, User := range mocks.User {
		if User.Uuid == id {
			mocks.User = append(mocks.User[:index], mocks.User[index+1:]...)

			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("deleted user")
			break
		}
	}
}
