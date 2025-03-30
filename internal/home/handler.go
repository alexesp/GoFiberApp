package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	//panic("panic")
	return c.SendString("Start")
}
func (h *HomeHandler) error(c *fiber.Ctx) error {
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
