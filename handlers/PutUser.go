package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/aelpxy/api-http/mocks"
	"github.com/aelpxy/api-http/models"
	"github.com/gorilla/mux"
)

func PutUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Panic(err)
	}

	var updatedUser models.User
	json.Unmarshal(body, &updatedUser)

	for index, User := range mocks.User {
		if User.Uuid == id {
			User.Username = updatedUser.Username
			User.Password = updatedUser.Password

			mocks.User[index] = User
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			json.NewEncoder(w).Encode("updated")
			break
		}
	}
}
