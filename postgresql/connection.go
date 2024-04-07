package configs

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
)

var (
	db    *sql.DB
	dbErr error
)

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
