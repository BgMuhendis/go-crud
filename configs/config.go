package configs

import (
	"os"
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