package handlers

import (
	"github.com/bekzod003/chat-api-gateway/genproto/auth_service"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var req auth_service.SignIpRequest
	err := c.BodyParser(&req)
	if err != nil {
		h.log.Error("Error whil parsing request from body to struct --->>>", err)
		return err
	}

	resp, err := h.services.AuthService().SignIn(
		c.Context(), &req,
	)
	if err != nil {
		h.log.Error("Error while signing in --->>>", err)
		return err
	}

	c.JSON(resp)
	return err
}

func (h *Handler) SignUp(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) Test(c *fiber.Ctx) error {
	print("Hello")
	c.JSON("Ping")
	return nil
}
