package handlers

import (
	"github.com/bekzod003/chat-api-gateway/config"
	"github.com/bekzod003/chat-api-gateway/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	cfg      config.Config
	log      logrus.Logger
	services services.ServiceManager
}

func NewHandler(cfg config.Config, log logrus.Logger, services services.ServiceManager) *Handler {
	return &Handler{
		cfg:      cfg,
		log:      log,
		services: services,
	}
}

func (h *Handler) handleResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	c.Context().SetStatusCode(statusCode)
	return c.JSON(data)
}
