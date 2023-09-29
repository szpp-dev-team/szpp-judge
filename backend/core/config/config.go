package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBUser    string `envconfig:"MYSQL_USER" required:"true"`
	DBPass    string `envconfig:"MYSQL_PASSWORD" required:"true"`
	DBAddr    string `envconfig:"BACKEND_DB_ADDR" required:"true"`
	DBName    string `envconfig:"MYSQL_DATABASE" required:"true"`
	GrpcPort  string `envconfig:"BACKEND_GRPC_PORT" default:"50051"`
	HttpPort  string `envconfig:"BACKEND_HTTP_PORT" default:"8080"`
	ModeDev   bool   `envconfig:"MODE_DEV" default:"true"`
	JudgeAddr string `envconfig:"BACKEND_JUDGE_ADDR" required:"true"`
}

func New() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	// for Cloud Run
	if port := os.Getenv("PORT"); port != "" {
		config.GrpcPort = port
	}
	return config, nil
}
