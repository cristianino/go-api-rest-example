package models

import (
	"fmt"
	"gomysql/db"
)

type User struct {
	Id       int64
	Name     string
	Username string
	Email    string
	Password string
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
func (user *User) Insert() {
	sql := "INSERT " + UserNameTable + " SET name=?, username=?, email=?, password=?"
	result, _ := db.Exec(sql, user.Name, user.Username, user.Email, user.Password)
	user.Id, _ = result.LastInsertId()
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
func ListUser() Users {
	query := "SELECT id, name, username, email FROM " + UserNameTable
	users := Users{}

	rows, _ := db.Query(query)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Email, &user.Name, &user.Username)
		users = append(users, user)
	}

	return users
}

// get user
func GetUser(id int) *User {
	user := User{}
	query := fmt.Sprintf("SELECT id, name, username, email FROM %s WHERE id=%d LIMIT 1000", UserNameTable, id)
	rows, _ := db.Query(query)

	for rows.Next() {
		rows.Scan(&user.Id, &user.Email, &user.Name, &user.Username)
	}
	return &user
}
