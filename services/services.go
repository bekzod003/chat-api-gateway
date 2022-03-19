package services

import (
	"github.com/bekzod003/chat-api-gateway/config"
	"github.com/bekzod003/chat-api-gateway/genproto/auth_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	// Implemetation of grpc services
	AuthService() auth_service.AuthServiceClient
}

type grpcClients struct {
	// grpc services
	authService auth_service.AuthServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManager, error) {
	connAuthService, err := grpc.Dial(
		cfg.AuthGrpcPort+cfg.AuthGrpcHost,
		grpc.WithInsecure(),
	)
	if err != nil {
		logrus.Error("ConnAuthService error --->>>", err)
		return nil, err
	}

	return &grpcClients{
		authService: auth_service.NewAuthServiceClient(connAuthService),
	}, nil
}

func (g *grpcClients) AuthService() auth_service.AuthServiceClient {
	return g.authService
}
