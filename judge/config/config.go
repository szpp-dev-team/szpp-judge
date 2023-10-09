package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	GrpcPort      string `envconfig:"JUDGE_GRPC_PORT" default:"50052"`
	ModeDev       bool   `envconfig:"MODE_DEV" default:"true"`
	WorkdirRoot   string `envconfig:"WORKDIR_ROOT" required:"true"`
	GcsBucketName string `envconfig:"GCS_BUCKET_NAME" required:"true"`
}

func New() (*Config, error) {
	c := &Config{}
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}
