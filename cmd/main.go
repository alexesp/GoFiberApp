package main

import (
	"log"

	"github.com/alexesp/FiberWebApp/config"
	"github.com/alexesp/FiberWebApp/internal/home"
	"github.com/gofiber/fiber/v2"
)

//func start(c *fiber.Ctx) error {
//return c.SendString("Start")
//}

func main() {
	//fmt.Println("Start")
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()
	//app.Get("/", start)
	home.NewHandler(app)
	app.Listen(":3000")
}
