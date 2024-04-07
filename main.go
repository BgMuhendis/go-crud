package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"go-crud/entity"
	"go-crud/postgresql"
	"go-crud/repository"
	"net/http"
	"strconv"
)

func main() {

	listen := make(chan int)

	app := fiber.New()

	resultConnect := configs.DBConnect()
	cityRepo := respository.NewRepo(resultConnect)

	api := app.Group("/city")
	api.Get("/", func(c fiber.Ctx) error {

		cityList := cityRepo.List()

		if len(cityList) == 0 {
			return nil
		}

		return c.Status(http.StatusOK).JSON(cityList)

	})

	api.Get("/:id", func(c fiber.Ctx) error {

		if queryId := c.Params("id"); queryId != "" {
			cityId, _ := strconv.Atoi(queryId)
			city := cityRepo.GetById(cityId)

			if city == nil {
				return c.Status(http.StatusNotFound).JSON("Not found")
			}

			return c.Status(http.StatusOK).JSON(city)

		}

		return nil

	})

	api.Post("/", func(c fiber.Ctx) error {
		var city entity.City

		bodyBytes := c.Body()

		if err := json.Unmarshal(bodyBytes, &city); err != nil {

			return c.Status(http.StatusBadRequest).JSON(err)

		}
		cityRepo.Insert(city)

		return nil

	})

	api.Delete("/:id", func(c fiber.Ctx) error {

		if queryId := c.Params("id"); queryId != "" {
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
