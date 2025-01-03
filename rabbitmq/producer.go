package rabbitmq

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type Producer struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queues  []string
}

func NewProducer(amqpURL string, queueNames []string, exchangeName string) (*Producer, error) {
	conn, err := amqp091.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	err = channel.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	for _, queue := range queueNames {
		_, err = channel.QueueDeclare(
			queue,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			_ = channel.Close()
			_ = conn.Close()
			return nil, err
		}
	}

	return &Producer{
		conn:    conn,
		channel: channel,
		queues:  queueNames,
	}, nil
}

func (p *Producer) Publish(ctx context.Context, message []byte, queue string) error {
	return p.channel.PublishWithContext(ctx,
		"",
		queue,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
}

func (p *Producer) Close() {
	if p.channel != nil {
		err := p.channel.Close()
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}
	if p.conn != nil {
		err := p.conn.Close()
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}
}
