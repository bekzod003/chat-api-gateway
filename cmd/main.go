package main

import (
	"github.com/bekzod003/chat-api-gateway/api"
	"github.com/bekzod003/chat-api-gateway/api/handlers"
	"github.com/bekzod003/chat-api-gateway/config"
	"github.com/bekzod003/chat-api-gateway/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("app is starting")

	cfg := config.Load()

	grpcServices, err := services.NewGrpcClients(cfg)
	if err != nil {
		logrus.Error("error while creating new grpc clients: ", err)
	}

	r := fiber.New(
		fiber.Config{
			AppName: cfg.ServiceName,
		},
	)

	h := handlers.NewHandler(cfg, logrus.Logger{}, grpcServices)
	api.SetUpApi(r, h, cfg)

	r.Listen(cfg.Port)
}
