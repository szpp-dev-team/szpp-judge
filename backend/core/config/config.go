package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBUser               string `envconfig:"MYSQL_USER" required:"true"`
	DBPass               string `envconfig:"MYSQL_PASSWORD" required:"true"`
	DBAddr               string `envconfig:"BACKEND_DB_ADDR" required:"true"`
	DBName               string `envconfig:"MYSQL_DATABASE" required:"true"`
	ConnectPort          string `envconfig:"BACKEND_CONNECT_PORT" default:"8080"`
	ModeDev              bool   `envconfig:"MODE_DEV" default:"true"`
	JudgeAddr            string `envconfig:"BACKEND_JUDGE_ADDR" required:"true"`
	JWTSecret            string `envconfig:"JWT_SECRET" required:"true"`
	CloudTasksProjectID  string `envconfig:"CLOUD_TASKS_PROJECT_ID" required:"true"`
	CloudTasksLocationID string `envconfig:"CLOUD_TASKS_LOCATION_ID" required:"true"`
	CloudTasksQueueID    string `envconfig:"CLOUD_TASKS_QUEUE_ID" required:"true"`
}

func New() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	// for Cloud Run
	if port := os.Getenv("PORT"); port != "" {
		config.ConnectPort = port
	}
	return config, nil
}
