package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const url = "goUser:goUser@tcp(localhost:3306)/goweb_db"

var db *sql.DB

// Open conection
func Connect() {
	var err error
	db, err = sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión exitosa.")
}

// Close conection
func Close() {
	db.Close()
}

// Create table
func CreateTable(schema string, nameTable string) {
	if !ExistsTable(nameTable) {
		_, err := db.Exec(schema)
		if err != nil {
			panic(err)
		}
	}
}

// Verificar si existe una tabla
func ExistsTable(tableName string) bool {
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	return rows.Next()
}

// Limpiar tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	_, err := Exec(sql)
	if err != nil {
		log.Println(err)
	}
}

// Verify conection
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	res, err := db.Exec(query, args...)

	if err != nil {
		log.Println(err)
	}
	return res, err
}

// Polimorfismo de Query
func Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil {
		log.Println(err)
	}
	return rows, err
}
