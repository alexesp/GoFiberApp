package main

import (
	"github.com/alexesp/FiberWebApp/internal/home"
	"github.com/gofiber/fiber/v2"
)

func start(c *fiber.Ctx) error {
	return c.SendString("Start")
}

func main() {
	//fmt.Println("Start")
	app := fiber.New()
	//app.Get("/", start)
	home.NewHandler(app)
	app.Listen(":3000")
}
