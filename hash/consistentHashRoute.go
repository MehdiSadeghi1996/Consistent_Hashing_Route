package hash

import (
	"crypto/sha256"
	"sort"
	"strconv"
)

type ConsistentHashRing struct {
	ring     map[int]string // Hash ring: maps hash values to queue names
	hashKeys []int          // Sorted hash keys for efficient lookup
	replicas int            // Number of virtual nodes per queue
}

func NewConsistentHashRing(replicas int) *ConsistentHashRing {
	return &ConsistentHashRing{
		ring:     make(map[int]string),
		hashKeys: []int{},
		replicas: replicas,
	}
}

// AddQueue adds a queue to the hash ring.
func (c *ConsistentHashRing) AddQueue(queue string) {
	for i := 0; i < c.replicas; i++ {
		// Create a virtual node key: "queue#replica"
		virtualNode := queue + "#" + strconv.Itoa(i)
		hash := hashKey(virtualNode)
		c.ring[hash] = queue
		c.hashKeys = append(c.hashKeys, hash)
	}
	sort.Ints(c.hashKeys) // Sort keys for efficient lookup
}

// RemoveQueue removes a queue from the hash ring.
func (c *ConsistentHashRing) RemoveQueue(queue string) {
	for i := 0; i < c.replicas; i++ {
		virtualNode := queue + "#" + strconv.Itoa(i)
		hash := hashKey(virtualNode)
		delete(c.ring, hash)
		c.removeHashKey(hash)
	}
}

// GetQueue gets the closest queue on the hash ring for the given hash key.
func (c *ConsistentHashRing) GetQueue(orderId int) string {
	if len(c.hashKeys) == 0 {
		return ""
	}
	hash := hashKey(strconv.Itoa(orderId))
	idx := c.searchClosest(hash)
	return c.ring[c.hashKeys[idx]]
}

// removeHashKey removes a hash key from the sorted list.
func (c *ConsistentHashRing) removeHashKey(hash int) {
	index := -1
	for i, v := range c.hashKeys {
		if v == hash {
			index = i
			break
		}
	}
	if index >= 0 {
		c.hashKeys = append(c.hashKeys[:index], c.hashKeys[index+1:]...)
	}
}

// searchClosest finds the index of the closest hash key on the ring.
func (c *ConsistentHashRing) searchClosest(hash int) int {
	idx := sort.Search(len(c.hashKeys), func(i int) bool {
		return c.hashKeys[i] >= hash
	})
	if idx == len(c.hashKeys) {
		return 0 // Wrap around to the first key
	}
	return idx
}

// hashKey generates a hash for a given key.
func hashKey(key string) int {
	hash := sha256.Sum256([]byte(key))
	return int(hash[0])<<24 | int(hash[1])<<16 | int(hash[2])<<8 | int(hash[3])
}
