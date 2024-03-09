package configs

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	db    *sql.DB
	dbErr error
)

type dbInfo struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func callDBInfo() *dbInfo {
	return &dbInfo{
		host:     os.Getenv("HOST"),
		port:     os.Getenv("PORT"),
		user:     os.Getenv("USER"),
		password: os.Getenv("PASSWORD"),
		dbname:   os.Getenv("DBNAME"),
	}
}

func DBConnect() *sql.DB {
	godotenv.Load(".env")
	connect := callDBInfo()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", connect.host, connect.port, connect.user, connect.password, connect.dbname)
	db, dbErr = sql.Open("postgres", psqlInfo)

	if dbErr != nil {
		panic(dbErr)
	}

	return db

}
