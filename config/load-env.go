package config

import (
	"github.com/joho/godotenv"
	"os"
)

type dbInfo struct {
	DBHOST     string
	DBPORT     string
	DBUSER     string
	DBPASSWORD string
	DBNAME     string
}

func init() {
	godotenv.Load(".env")
}

func CallDBInfo() *dbInfo {
	return &dbInfo{
		DBHOST:     os.Getenv("HOST"),
		DBPORT:     os.Getenv("PORT"),
		DBUSER:     os.Getenv("USER"),
		DBPASSWORD: os.Getenv("PASSWORD"),
		DBNAME:     os.Getenv("DBNAME"),
	}
}
