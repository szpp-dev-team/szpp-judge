package handlejudgetask

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBUser    string `envconfig:"MYSQL_USER" required:"true"`
	DBPass    string `envconfig:"MYSQL_PASSWORD" required:"true"`
	DBAddr    string `envconfig:"BACKEND_DB_ADDR" required:"true"`
	DBName    string `envconfig:"MYSQL_DATABASE" required:"true"`
	JudgeAddr string `envconfig:"BACKEND_JUDGE_ADDR" required:"true"`
}

func newConfig() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
