package handlers

import (
	"encoding/json"
	"fmt"
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
	}
	output, _ := json.Marshal(users)
	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	users, err := models.GetUser(userID)
	if err != nil {
		log.Println(err)
	}
	output, _ := json.Marshal(users)

	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(rw, string(output))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
	} else {
		user.Save()
	}
	output, _ := json.Marshal(user)
	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(rw, string(output))
}

func EditUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
	} else {
		user.Id = int64(userID)
		user.Update()
	}
	output, _ := json.Marshal(user)
	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(rw, string(output))
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	users, err := models.GetUser(userID)
	if err != nil {
		log.Println(err)
	}
	users.Delete()

	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(rw, "")
}
