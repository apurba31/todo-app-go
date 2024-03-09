package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type Todo struct {
	ID    int    `json: "id"`
	Title string `json: "title"`
	Done  bool   `json: "done"`
	Body  string `json: "body"`
}

func main() {
	fmt.Println("Hello")

	app := fiber.New()

	todos := []Todo{}

	app.Get("/healthCheck", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c fiber.Ctx) error {
		todo := &Todo{}

		body := c.Body()

		error := json.Unmarshal(body, todo)
		if error != nil {
			return c.Status(http.StatusBadRequest).SendString("Error decoding JSON body")
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("/api/todos/", func(c fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}
