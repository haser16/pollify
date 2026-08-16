package users_postgres_repository

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AuthorizeConfig struct {
	SECRET string `envconfig:"SECRET"`
}

func NewConfig() (AuthorizeConfig, error) {
	var config AuthorizeConfig

	if err := envconfig.Process("JWT", &config); err != nil {
		return AuthorizeConfig{}, fmt.Errorf("process env config: %w", err)
	}

	return config, nil
}

func NewConfigMust() AuthorizeConfig {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get JWT secret key: %w", err)
		panic(err)
	}
	return config
}
