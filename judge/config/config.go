package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	GrpcPort string `envconfig:"JUDGE_GRPC_PORT" default:"50052"`
	ModeDev  bool   `envconfig:"MODE_DEV" default:"true"`
}

type DevConfig struct {
	WorkingDirAbsRoot string `envconfig:"JUDGE_DEV_WORKING_DIR_ABS_ROOT" required:"true"`
}

func New() (*Config, error) {
	c := &Config{}
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}

func NewDevConfig() (*DevConfig, error) {
	c := &DevConfig{}
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}
