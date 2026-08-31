package core_publisher_rabbitmq

import (
	"context"
	"encoding/json"

	core_publisher "pollify/internal/core/publisher"

	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   *amqp091.Queue
}

func NewConsumer(url, queueName string) (*Consumer, error) {
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

	return &Consumer{
		conn:    conn,
		channel: ch,
		queue:   &queue,
	}, nil
}

func (c *Consumer) Consume(
	ctx context.Context,
	handler func(core_publisher.VerificationMessage) error,
) error {
	msgs, err := c.channel.Consume(
		c.queue.Name,
		"",
		false, // autoAck = false
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case msg, ok := <-msgs:
			if !ok {
				return nil
			}

			var message core_publisher.VerificationMessage

			if err := json.Unmarshal(msg.Body, &message); err != nil {
				// Сообщение битое — повторять его бессмысленно.
				_ = msg.Nack(false, false)
				continue
			}

			if err := handler(message); err != nil {
				// Email не отправился — возвращаем сообщение в очередь.
				_ = msg.Nack(false, true)
				continue
			}

			// Email успешно отправлен.
			if err := msg.Ack(false); err != nil {
				return err
			}
		}
	}
}

func (c *Consumer) Close() {
	_ = c.channel.Close()
	_ = c.conn.Close()
}
