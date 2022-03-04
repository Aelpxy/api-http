package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/aelpxy/api-http/mocks"
	"github.com/aelpxy/api-http/models"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Panic(err)
	}

	var User models.User
	json.Unmarshal(body, &User)

	User.Uuid = rand.Intn(100)

	mocks.User = append(mocks.User, User)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("created")
}
