package models

import (
	"fmt"

	"github.com/cristianino/go-api-rest-example/db"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

const UserSchema string = `CREATE TABLE ` + UserNameTable + ` (
	id INT(6) UNSIGNED AUTO_INCREMENT,
	name VARCHAR(30) NOT NULL,
	username VARCHAR(30) NOT NULL,
	email VARCHAR(50) NOT NULL,
	password VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(id)
);`

const UserNameTable string = "users"

func (user *User) NewUser(name, username, email, password string) {
	user.Email = email
	user.Name = name
	user.Password = password
	user.Username = username
}

func CreateUser(name, username, email, password string) *User {
	user := User{
		Name:     name,
		Username: username,
		Email:    email,
		Password: password,
	}
	user.Save()
	return &user
}

// Insertar Registro
func (user *User) Insert() error {
	sql := "INSERT " + UserNameTable + " SET name=?, username=?, email=?, password=?"
	result, err := db.Exec(sql, user.Name, user.Username, user.Email, user.Password)

	if err != nil {
		return err
	}
	user.Id, _ = result.LastInsertId()
	return nil
}

func (user *User) Save() {
	if user.Id == 0 {
		user.Insert()
	} else {
		user.Update()
	}
}

// update Registro
func (user *User) Update() {
	sql := "UPDATE " + UserNameTable + " SET name=?, username=?, email=?, password=? WHERE id=?"
	db.Exec(sql, user.Name, user.Username, user.Email, user.Password, user.Id)
}

// delete Registro
func (user *User) Delete() {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id=%d", UserNameTable, user.Id)
	db.Exec(sql)
}

// List
func ListUser() (Users, error) {
	query := "SELECT id, name, username, email FROM " + UserNameTable
	users := Users{}

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Email, &user.Name, &user.Username)
		users = append(users, user)
	}

	return users, nil
}

// get user
func GetUser(id int) (*User, error) {
	user := User{}
	query := fmt.Sprintf("SELECT id, name, username, email FROM %s WHERE id=%d LIMIT 1000", UserNameTable, id)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Email, &user.Name, &user.Username)
	}
	return &user, nil
}
