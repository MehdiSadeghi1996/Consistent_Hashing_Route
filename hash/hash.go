package hash

import (
	"log"
	"strconv"
)

type HashingRoute struct {
}

func (*HashingRoute) GetRouteByOrderId(orderId int, queues []string) string {
	index := orderId % len(queues)
	return queues[index]
}

func (b *HashingRoute) GetAssignedQueue(instanceIndex string, queues []string) string {
	if len(queues) == 0 || instanceIndex == "" {
		log.Fatalf("queues or instanceIndex not set")
	}

	number, err := strconv.Atoi(instanceIndex)
	if err != nil || number < 0 || number >= len(queues) {
		log.Fatalf("Invalid instance Specific Index: %v", number)
	}

	return queues[number]
}
