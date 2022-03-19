package api

import (
	"github.com/bekzod003/chat-api-gateway/api/handlers"
	"github.com/bekzod003/chat-api-gateway/config"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetUpApi(r *fiber.App, h *handlers.Handler, cfg config.Config) {
	logrus.Info("Api is being set up")
	r.Get("/ping", h.Test)

	auth := r.Group("/auth")
	{
		auth.Post("/sign-up", h.SignUp)
		auth.Post("/sign-in", h.SignIn)
	}
}
