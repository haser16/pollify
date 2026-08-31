package main

import (
	"context"
	"fmt"
	"log"
	"os"
	core_publisher "pollify/internal/core/publisher"

	core_publisher_rabbitmq "pollify/internal/core/publisher/rabbitmq"

	email_config "pollify/services/email-deliver/config"
	"pollify/services/email-deliver/email"
	services_email_logger "pollify/services/email-deliver/logger"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}

	config := email_config.NewConfigMust()

	logger, err := services_email_logger.NewLogger(
		services_email_logger.NewConfigMust(),
	)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %s\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	sender, err := email.NewSender(
		config.SMTPHost,
		config.SMTPPort,
		config.SMTPUser,
		config.SMTPPassword,
		config.SMTPFrom,
	)
	if err != nil {
		logger.Fatal(
			"Failed to initialize email sender",
			zap.Error(err),
		)
	}

	consumer, err := core_publisher_rabbitmq.NewConsumer(
		config.PublisherURL,
		config.QueueEmailName,
	)
	if err != nil {
		logger.Fatal(
			"Failed to initialize rabbitmq consumer",
			zap.Error(err),
		)
	}
	defer consumer.Close()

	logger.Info("start email-deliver service")

	ctx := context.Background()

	err = consumer.Consume(ctx, func(
		message core_publisher.VerificationMessage,
	) error {
		logger.Debug(
			"Received email message",
			zap.String("token", message.Token),
		)

		return sender.SendVerificationEmail(
			message.Email,
			message.Token,
		)
	})

	if err != nil {
		logger.Fatal(
			"Email consumer stopped",
			zap.Error(err),
		)
	}
}
