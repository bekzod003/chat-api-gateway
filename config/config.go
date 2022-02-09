package config

type Config struct {
	AuthGrpcPort string
	AuthGrpcHost string
}

func Load() Config {
	return Config{
		AuthGrpcPort: "",
		AuthGrpcHost: "",
	}
}
