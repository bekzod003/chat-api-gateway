package main

import (
	"github.com/bekzod003/chat-api-gateway/api/handlers"
	"github.com/bekzod003/chat-api-gateway/config"
	"github.com/bekzod003/chat-api-gateway/services"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("app is starting")

	cfg := config.Load()

	grpcServices, err := services.NewGrpcClients(cfg)
	if err != nil {
		logrus.Error("error while creating new grpc clients: ", err)
	}

	handler := handlers.NewHandler(cfg, *logrus.New(), grpcServices)
	app := handler.InitRoutes()

	app.Listen(cfg.Port)
}
