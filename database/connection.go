package database

import (
	"database/sql"
	"fmt"
	"go-crud/config"
	_ "github.com/lib/pq"
)

var (
	db    *sql.DB
	dbErr error
)

func DBConnect() *sql.DB {
	connect := config.CallDBInfo()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", connect.DBHOST, connect.DBPORT, connect.DBUSER, connect.DBPASSWORD, connect.DBNAME)
	db, dbErr = sql.Open("postgres", psqlInfo)

	if dbErr != nil {
		panic(dbErr)
	}

	return db

}
