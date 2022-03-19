package api

import (
	"github.com/bekzod003/chat-api-gateway/api/handlers"
	"github.com/bekzod003/chat-api-gateway/config"
	"github.com/gofiber/fiber/v2"
)

func SetUpApi(r *fiber.App, h *handlers.Handler, cfg config.Config) {
	println("Api is being set up")
	r.Get("/ping", h.Test)

	// auth := r.Group("/auth")
	// auth.Use(func(c *fiber.Ctx) error {
	// 	return c.Next()
	// })
	// {
	// 	auth.Post("/sing-up", h.SignUp)
	// 	auth.Post("sing-in", h.SignIn)
	// }
	r.Post("/auth/sign-up", h.SignIn)

}
