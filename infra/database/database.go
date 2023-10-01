package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "tugasaden"
)

var (
	db  *sql.DB
	err error
)

func handlerDatabaseConnection() {
	sqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//psqlinfo := "host=localhost port=5432 user=postgres password=postgres dbname=tugasaden sslmode=disable"

	db, err = sql.Open("postgres", sqlinfo)

	if err != nil {

		log.Panic("Error saat validasi database argumen", err)
	}

	err = db.Ping()

	if err != nil {

		log.Panic("Error saat koneksi ke database", err)
	}
}

func InitDatabase() {
	handlerDatabaseConnection()
}

func GetDatabaseInstance() *sql.DB {

	if db == nil {
		log.Panic("Erorr saat get db")
	}
	return db
}
