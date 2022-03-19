package handlers

import (
	"github.com/bekzod003/chat-api-gateway/config"
	"github.com/bekzod003/chat-api-gateway/services"
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
