package handlers

import (
	"net/http"

	"github.com/bekzod003/chat-api-gateway/genproto/auth_service"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	req := new(auth_service.SignUpRequest)
	err := c.BodyParser(req)
	if err != nil {
		h.log.Error("Error while parsing from body request --->>> ", err)
		return h.handleResponse(c, http.StatusBadRequest, err)
	}

	res, err := h.services.AuthService().SignUp(
		c.Context(),
		req,
	)
	if err != nil {
		h.log.Error("Grpc error in auth service --->>> ", err)
		return h.handleResponse(c, http.StatusInternalServerError, err)
	}

	return h.handleResponse(c, http.StatusCreated, res)
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	return nil
}
