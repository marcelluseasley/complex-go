package main

import (


	"strconv"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func main() {




	// Fiber route handlers

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("hello world")
	})

	app.Get("/values/all", func(c *fiber.Ctx) {
		//get all values from postgress
		c.Send("all the values")
	})

	app.Post("/values", func(c *fiber.Ctx) {
		index := c.FormValue("index")
		i, _ := strconv.ParseInt(index, 10, 32)

		if i > 40 {
			c.Status(422).Send("Index too high")
		}

		
		c.Send(fib(i))
	})

	app.Listen(5000)

}

func fib(i int64) int64 {
	if i < 2 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}
