package core_config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string        `envconfig:"POSTGRES_HOST" required:"true"`
	Port     string        `envconfig:"POSTGRES_PORT" default:"5432"`
	User     string        `envconfig:"POSTGRES_USER" required:"true"`
	Password string        `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Database string        `envconfig:"POSTGRES_DB" required:"true"`
	TimeOut  time.Duration `envconfig:"POSTGRES_TIMEOUT" required:"true"`

	JWTToken string `envconfig:"JWT_TOKEN" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process env config: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get config: %w", err)
		panic(err)
	}
	return config
}
