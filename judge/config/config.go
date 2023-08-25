package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	GrpcPort          string `envconfig:"JUDGE_GRPC_PORT" default:"50052"`
	WorkingDirAbsRoot string `envconfig:"JUDGE_WORKING_DIR_ABS_ROOT"`
	ModeDev           bool   `envconfig:"MODE_DEV" default:"true"`
}

func New() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
