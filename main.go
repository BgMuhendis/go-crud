package main

import (
	"encoding/json"
<<<<<<< HEAD
	"github.com/gofiber/fiber/v3"
=======
>>>>>>> 95525087c828c83aa5beb488544808bb1a855014
	"go-crud/entity"
	"go-crud/postgresql"
	"go-crud/repository"
	"net/http"
	"strconv"
	"github.com/gofiber/fiber/v3"
)

func main() {

	listen := make(chan int)

	app := fiber.New()

	resultConnect := configs.DBConnect()
	cityRepo := respository.NewRepo(resultConnect)

<<<<<<< HEAD
	api := app.Group("/city")
	api.Get("/", func(c fiber.Ctx) error {

		cityList := cityRepo.List()

		if len(cityList) == 0 {
=======

	api := app.Group("/city")
	api.Get("/", func(c fiber.Ctx) error {

		cityList := cityRepo.List()

		if len(cityList) ==0 {
>>>>>>> 95525087c828c83aa5beb488544808bb1a855014
			return nil
		}
		
		return c.Status(http.StatusOK).JSON(cityList)


		return c.Status(http.StatusOK).JSON(cityList)

	})

	api.Get("/:id", func(c fiber.Ctx) error {

<<<<<<< HEAD
		if queryId := c.Params("id"); queryId != "" {
=======
		if queryId := c.Params("id"); queryId !="" {
>>>>>>> 95525087c828c83aa5beb488544808bb1a855014
			cityId, _ := strconv.Atoi(queryId)
			city := cityRepo.GetById(cityId)

			if city == nil {
				return c.Status(http.StatusNotFound).JSON("Not found")
			}

			return c.Status(http.StatusOK).JSON(city)
<<<<<<< HEAD

=======
			
>>>>>>> 95525087c828c83aa5beb488544808bb1a855014
		}

		return nil

	})

	api.Post("/", func(c fiber.Ctx) error {
		var city entity.City

<<<<<<< HEAD
		bodyBytes := c.Body()
=======
		bodyBytes:= c.Body()
>>>>>>> 95525087c828c83aa5beb488544808bb1a855014

		if err := json.Unmarshal(bodyBytes, &city); err != nil {

			return c.Status(http.StatusBadRequest).JSON(err)
<<<<<<< HEAD

=======
		
>>>>>>> 95525087c828c83aa5beb488544808bb1a855014
		}
		cityRepo.Insert(city)

		return nil

	})

	api.Delete("/:id", func(c fiber.Ctx) error {

<<<<<<< HEAD
		if queryId := c.Params("id"); queryId != "" {
			cityId, _ := strconv.Atoi(queryId)
			cityRepo.DeleteById(cityId)

=======
		if queryId := c.Params("id"); queryId !="" {
			cityId, _ := strconv.Atoi(queryId)
			cityRepo.DeleteById(cityId)
			
>>>>>>> 95525087c828c83aa5beb488544808bb1a855014
		}
		return nil

	})

<<<<<<< HEAD
	go func() {
		app.Listen(":3000")

		listen <- 1

	}()

	<-listen
=======

	go func ()  {
		app.Listen(":3000")

		listen <-1

	}()	

	<- listen
>>>>>>> 95525087c828c83aa5beb488544808bb1a855014

}
