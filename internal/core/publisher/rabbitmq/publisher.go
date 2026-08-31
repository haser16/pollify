package core_publisher_rabbitmq

import (
	"context"
	"encoding/json"
	core_publisher "pollify/internal/core/publisher"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   *amqp091.Queue
}

func NewPublisher(url, queueName string) (*Publisher, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	queue, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		true,
		amqp091.Table{
			amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
		},
	)
	if err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, err
	}

	return &Publisher{
		conn:    conn,
		channel: ch,
		queue:   &queue,
	}, nil
}

func (p *Publisher) Publish(
	ctx context.Context,
	message core_publisher.VerificationMessage,
) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.channel.PublishWithContext(
		ctx,
		"",
		p.queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (p *Publisher) Close() {
	_ = p.channel.Close()
	_ = p.conn.Close()
}
