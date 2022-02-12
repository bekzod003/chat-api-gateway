package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	// Port of the api gateway
	Port        string
	ServiceName string

	AuthGrpcPort string
	AuthGrpcHost string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		logrus.Info("No .env file was found")
	}

	return Config{
		Port:         getOrReturnDefaultValue("PORT", ":8000"),
		AuthGrpcPort: getOrReturnDefaultValue("PORT", ":8001"),
		AuthGrpcHost: getOrReturnDefaultValue("AUTH_GRPC_PORT", "localhost"),
	}
}

// Sets value from env variable, if not exists set default value
func getOrReturnDefaultValue(key string, defaultValue interface{}) string {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return fmt.Sprint(defaultValue)
}
