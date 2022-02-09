package services

import "github.com/bekzod003/chat-api-gateway/config"

type ServiceManager interface {
	// Implemetation of grpc services
}

type grpcClients struct {
	// grpc services
}

func NewGrpcClients(cfg config.Config) (ServiceManager, error) {
	return nil, nil
}
