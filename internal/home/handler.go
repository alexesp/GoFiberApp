package home

import (
	"github.com/alexesp/FiberWebApp/pkg/tadapter"
	"github.com/alexesp/FiberWebApp/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type HomeHandler struct {
	router fiber.Router
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	//panic("panic")
	//tmpl, err := template.New("test").Parse("{{.Count}} - numero de usuarios")
	//tmpl := template.Must(template.ParseFiles("./html/page.html"))
	//data := struct{ Count int }{Count: 1}
	//if err != nil {
	//return fiber.NewError(fiber.StatusBadRequest, "Template error")
	//}
	// var tpl bytes.Buffer
	// if err := tmpl.Execute(&tpl, data); err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, "Template error")
	// }
	// c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	// return c.Send(tpl.Bytes())
	//return c.SendString("Start")
	// data := struct {
	// 	Count   int
	// 	IsAdmin bool
	// }{
	// 	Count:   10,
	// 	IsAdmin: false,
	// }
	//return c.Render("page", data)
	// return c.Render("page", fiber.Map{
	// 	"Count": 5,
	// })
	component := views.Hello("Hello Fiber")
	//component.Render(context.Background(), os.Stdout)

	//return c.SendString("Test template")
	return tadapter.Render(c, component)
}
func (h *HomeHandler) error(c *fiber.Ctx) error {
	log.Info("Info")
	log.Debug("Debug")
	log.Warn("Warn")
	log.Error("Error")
	log.Panic("Error")
	return c.SendString("Error")
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}

	api := h.router.Group("/api")

	//h.router.Get("/", h.home)
	api.Get("/", h.home)
	api.Get("/error", h.error)
}
