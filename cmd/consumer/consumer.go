package main

import (
	"ConsistentHashing/config"
	"ConsistentHashing/dto"
	"ConsistentHashing/hash"
	"ConsistentHashing/rabbitmq"
	"context"
	"encoding/json"
	"log"
)

func main() {

	config.LoadConfig()

	amqpURL := config.GetRabbitMQURL()
	queueNames := config.GetQueueName()

	instanceIndexStr := config.GetNodeName()

	hf := hash.HashingRoute{}
	targetQueue := hf.GetAssignedQueue(instanceIndexStr, queueNames)

	consumer, err := rabbitmq.NewConsumer(amqpURL, targetQueue)
	if err != nil {
		log.Fatalf("Failed to initialize consumer: %v", err)
	}
	defer consumer.Close()

	handler := func(body []byte) error {
		var order dto.Order
		if err := json.Unmarshal(body, &order); err != nil {
			log.Printf("Failed to deserialize message: %v", err)
			return err
		}

		log.Printf("Processing order: %+v", order)
		return nil
	}

	err = consumer.StartConsuming(handler)
	if err != nil {
		log.Fatalf("Failed to start consuming messages: %v", err)
	}

	log.Println("Consumer is running. Press CTRL+C to exit.")
	ctx := context.Background()
	select {
	case <-ctx.Done():
		log.Println("Shutting down consumer...")
	}
}
