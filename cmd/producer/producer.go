package main

import (
	"ConsistentHashing/config"
	"ConsistentHashing/dto"
	"ConsistentHashing/hash"
	"ConsistentHashing/rabbitmq"
	"context"
	"encoding/json"
	"log"
	"time"
)

func main() {
	startProducer()
}

func startProducer() {
	produceMessages()
	log.Println("All messages have been sent.")
	log.Println("Shutting down producer...")
}

func produceMessages() {

	config.LoadConfig()
	amqpURL := config.GetRabbitMQURL()
	queueNames := config.GetQueueName()
	exchangeName := config.GetExchangeName()

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	producer, err := rabbitmq.NewProducer(amqpURL, queueNames, exchangeName)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ producer: %v", err)
	}
	defer producer.Close()

	for i := 0; i < 10000; i++ {
		orders := createOrders(i)
		for _, order := range orders {
			if err := produceToQueue(order, ctx, producer, queueNames); err != nil {
				log.Fatalln(err.Error())
				return
			}
		}
	}
}

func createOrders(id int) []*dto.Order {
	return []*dto.Order{
		{Id: id, IsPaid: id%10 == 0, Type: "X"},
		{Id: id, IsPaid: id%10 == 0, Type: "Y"},
		{Id: id, IsPaid: id%10 == 0, Type: "W"},
		{Id: id, IsPaid: id%10 == 0, Type: "Z"},
	}
}

func produceToQueue(order *dto.Order, ctx context.Context, producer *rabbitmq.Producer, queues []string) error {
	messageBody, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Failed to serialize order: %v", err)
		return err
	}

	hf := hash.HashingRoute{}
	targetQueue := hf.GetRouteByOrderId(order.Id, queues)

	if err := producer.Publish(ctx, messageBody, targetQueue); err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	log.Printf("Message sent: %+v to queue: %s", order, targetQueue)
	return nil
}
