package services_email_config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PublisherURL   string `envconfig:"PUBLISHER_URL" required:"true"`
	QueueEmailName string `envconfig:"QUEUE_EMAIL_NAME" required:"true"`
	SMTPHost       string `envconfig:"SMTP_HOST" required:"true"`
	SMTPPort       int    `envconfig:"SMTP_PORT" required:"true"`
	SMTPUser       string `envconfig:"SMTP_USER" required:"true"`
	SMTPPassword   string `envconfig:"SMTP_PASSWORD" required:"true"`
	SMTPFrom       string `envconfig:"SMTP_FROM" required:"true"`
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
