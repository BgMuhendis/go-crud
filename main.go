package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"net/http"
	"go-crud/repository"
	"go-crud/entity"
)

var (
	db    *sql.DB
	dbErr error
)

func main() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "159753+-+-"
	dbName := "godb"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, dbErr = sql.Open("postgres", psqlInfo)
	if dbErr != nil {
		panic(dbErr)
	}

	cityRepo := respository.NewRepo(db)

	http.HandleFunc("/city", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			cityList := cityRepo.List()
			(json.NewEncoder(writer).Encode(cityList))

		} else if request.Method == http.MethodPost {
			var city entity.City
			bodyBytes, err := io.ReadAll(request.Body)

			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
			}

			if err := json.Unmarshal(bodyBytes, &city); err != nil {

				http.Error(writer, err.Error(), http.StatusBadRequest)

				return

			}
			cityRepo.Insert(city)
			writer.WriteHeader(http.StatusCreated)
		} else {
			http.Error(writer, "Unsupported http method", http.StatusMethodNotAllowed)
			return
		}

	})

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		fmt.Println(err)
	}

}
