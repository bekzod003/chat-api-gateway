package handlers

import (
	"net/http"

	"github.com/bekzod003/chat-api-gateway/genproto/auth_service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	req := new(auth_service.SignUpRequest)
	err := c.BodyParser(req)
	if err != nil {
		logrus.Error("Error while parsing from body request --->>> ", err)
		c.JSON(err)
		return err
	}

	res, err := h.services.AuthService().SignUp(
		c.Context(),
		req,
	)
	if err != nil {
		logrus.Error("Grpc error in auth service --->>> ", err)
		c.JSON(err)
		return err
	}

	return h.handleResponse(c, http.StatusCreated, res)
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) Test(c *fiber.Ctx) error {
	logrus.Error("all")
	println("Hello")
	c.JSON("Ping")
	return nil
}
