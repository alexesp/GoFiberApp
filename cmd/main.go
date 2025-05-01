package main

import (
	"github.com/alexesp/FiberWebApp/config"
	"github.com/alexesp/FiberWebApp/internal/home"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

//func start(c *fiber.Ctx) error {
//return c.SendString("Start")
//}

func main() {
	//fmt.Println("Start")
	config.Init()
	config.NewDatabaseConfig()

	logConfig := config.NewLogConfig()
	//dbConf := config.NewDatabaseConfig()
	//log.Println(dbConf)

	//log.SetLevel(log.LevelTrace)
	log.SetLevel(log.Level(logConfig.Level))
	engine := html.New("./html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(recover.New())
	//app.Get("/", start)
	home.NewHandler(app)
	app.Listen(":3000")
}
