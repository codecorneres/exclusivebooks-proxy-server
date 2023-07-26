package main

import (
	"github.com/codecorneres/exclusivebooks-proxy-server/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/fanatics-join", routes.JoinFanaticsHandler)
	app.Post("/merge-fanatics-customer", routes.MergeFanaticsCustomerHandler)

	app.Listen(":3000")

}
