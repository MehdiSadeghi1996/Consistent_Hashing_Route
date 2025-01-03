package rabbitmq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type MessageHandler func(body []byte) error

type Consumer struct {
	conn      *amqp091.Connection
	channel   *amqp091.Channel
	queueName string
}

func NewConsumer(rabbitURL string, queueName string) (*Consumer, error) {
	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.Qos(
		1,
		0,
		false,
	)

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:      conn,
		channel:   ch,
		queueName: queueName,
	}, nil
}

func (c *Consumer) StartConsuming(handler MessageHandler) error {
	msgs, err := c.channel.Consume(
		c.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			err := handler(d.Body)
			if err != nil {
				log.Printf("Error processing message: %v", err)
			}
		}
	}()

	return nil
}

func (c *Consumer) Close() {
	if c.channel != nil {
		_ = c.channel.Close()
	}
	if c.conn != nil {
		_ = c.conn.Close()
	}
}
