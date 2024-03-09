package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"go-crud/configs"
	"go-crud/entity"
	"go-crud/repository"
	"io"
	"net/http"
	"strconv"
)

func main() {

	resultConnect := configs.DBConnect()
	cityRepo := respository.NewRepo(resultConnect)

	http.HandleFunc("/city", func(writer http.ResponseWriter, request *http.Request) {

		switch request.Method {

		case http.MethodGet:

			if request.URL.Query().Has("id") {
				queryId := request.URL.Query().Get("id")
				cityId, _ := strconv.Atoi(queryId)
				city := cityRepo.GetById(cityId)

				if city == nil {
					writer.WriteHeader(http.StatusNotFound)
					return
				}

				cityBytes, _ := json.Marshal(city)
				writer.Write(cityBytes)
				return
			}
			cityList := cityRepo.List()
			(json.NewEncoder(writer).Encode(cityList))

		case http.MethodPost:
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

		case http.MethodDelete:
			if request.URL.Query().Has("id") {
				queryId := request.URL.Query().Get("id")
				cityId, _ := strconv.Atoi(queryId)
				cityRepo.DeleteById(cityId)
			}

		default:
			http.Error(writer, "Unsupported http method", http.StatusMethodNotAllowed)
			return

		}

	})

	err := http.ListenAndServe("localhost:3000", nil)

	if err != nil {
		fmt.Println(err)
	}

}
