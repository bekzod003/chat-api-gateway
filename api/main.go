package api

import (
	"github.com/bekzod003/chat-api-gateway/api/handlers"
	"github.com/bekzod003/chat-api-gateway/config"
	"github.com/gofiber/fiber/v2"
)

func SetUpApi(r *fiber.App, h *handlers.Handler, cfg config.Config) {
	auth := r.Group("/auth")
	{
		auth.Post("/sign-up", h.SignUp)
		auth.Post("/sign-in", h.SignIn)
	}
}
