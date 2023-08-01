package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cristianino/go-api-rest-example/models"
	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users, err := models.ListUser()
	if err != nil {
		log.Println(err)
		models.SendNotFound(rw)
	}
	models.SendData(rw, users)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	users, err := models.GetUser(userID)
	if err != nil {
		log.Println(err)
		models.SendNotFound(rw)
	}
	models.SendData(rw, users)
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
		models.SendUnprocessableEntity(rw)
	}
	user.Save()
	models.SendData(rw, user)
}

func EditUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
		models.SendUnprocessableEntity(rw)
	}
	user.Id = int64(userID)
	user.Update()

	models.SendData(rw, user)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	users, err := models.GetUser(userID)
	if err != nil {
		log.Println(err)
		models.SendNotFound(rw)
	}
	users.Delete()

	models.SendData(rw, "")
}
