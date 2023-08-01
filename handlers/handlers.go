package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cristianino/go-api-rest-example/db"
	"github.com/cristianino/go-api-rest-example/models"
	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	db.Connect()
	defer db.Close()
	users := models.ListUser()
	output, _ := json.Marshal(users)
	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	db.Connect()
	defer db.Close()

	users := models.GetUser(userID)
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
		db.Connect()
		defer db.Close()
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
		db.Connect()
		defer db.Close()
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

	db.Connect()
	defer db.Close()

	users := models.GetUser(userID)
	users.Delete()

	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(rw, "")
}
