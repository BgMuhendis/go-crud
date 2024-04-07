package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
	"go-crud/database"
	"go-crud/model"
	"go-crud/repository"
	"net/http"
	"strconv"
)

func main() {

	listen := make(chan int)

	app := fiber.New()

	resultConnect := database.DBConnect()
	cityRepo := respository.NewRepo(resultConnect)

	api := app.Group("/city")
	api.Get("/", func(ctx fiber.Ctx) error {

		cityList := cityRepo.List()

		if len(cityList) == 0 {
			return nil
		}

		return ctx.Status(http.StatusOK).JSON(cityList)

	})

	api.Get("/:id", func(ctx fiber.Ctx) error {

		if queryId := ctx.Params("id"); queryId != "" {
			cityId, _ := strconv.Atoi(queryId)
			city := cityRepo.GetById(cityId)

			if city == nil {
				return ctx.Status(http.StatusNotFound).JSON("Not found")
			}

			return ctx.Status(http.StatusOK).JSON(city)

		}

		return nil

	})

	api.Post("/", func(ctx fiber.Ctx) error {
		var city model.City

		bodyBytes := ctx.Body()

		if err := json.Unmarshal(bodyBytes, &city); err != nil {

			return ctx.Status(http.StatusBadRequest).JSON(err)

		}
		cityRepo.Insert(city)

		return nil

	})

	api.Delete("/:id", func(ctx fiber.Ctx) error {

		if queryId := ctx.Params("id"); queryId != "" {
			cityId, _ := strconv.Atoi(queryId)
			cityRepo.DeleteById(cityId)

		}
		return nil

	})
	go func() {

		app.Listen(":3000")

		listen <- 1

	}()

	<-listen

}
