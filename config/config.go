package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func LoadConfig() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("No .env file found, loading defaults.")
	}

	viper.SetDefault("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	viper.SetDefault("RABBITMQ_QUEUE", "default_queue")
	viper.SetDefault("RABBITMQ_EXCHANGE", "default_exchange")
	viper.AutomaticEnv()
}

func GetRabbitMQURL() string {
	return viper.GetString("RABBITMQ_URL")
}

func GetQueueName() []string {
	queues := viper.GetString("RABBITMQ_QUEUE")
	return strings.Split(queues, ",")
}

func GetExchangeName() string {
	return viper.GetString("RABBITMQ_EXCHANGE")
}

func GetNodeName() string {
	return viper.GetString("INSTANCE_INDEX")
}
