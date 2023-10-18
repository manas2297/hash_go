package hashmap

import "fmt"

const (
	INITIAL_SIZE     = 1 << 4
	MAXIMUM_CAPACITY = 1 << 30
)

type HashMap struct {
	table    []*Entry
	size     int
	capacity int
}

func NewHashMap(capacity int) *HashMap {
	if capacity <= 0 || capacity > MAXIMUM_CAPACITY {
		capacity = INITIAL_SIZE
	} else {
		capacity = tableSizeFor(capacity)
	}
	return &HashMap{
		table:    make([]*Entry, capacity),
		capacity: 16,
	}
}

func tableSizeFor(cap int) int {
	n := cap - 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	if n < 0 {
		return 1
	} else {
		if n >= MAXIMUM_CAPACITY {
			return MAXIMUM_CAPACITY
		}
		return n + 1
	}
}

func (hm *HashMap) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash + int(char)) % hm.capacity
	}
	return hash
}

func (hm *HashMap) Set(key interface{}, value interface{}) {
	index := hm.hash(fmt.Sprintf("%v", key)) // Convert the key to a string for hashing
	node := hm.table[index]
	if node == nil {
		hm.table[index] = NewEntry(key, value)
	} else {
		previous := node
		for node != nil {
			if node.Key == key {
				node.Value = value
				return
			}
			previous = node
			node = node.Next
		}
		previous.Next = NewEntry(key, value)
	}
	hm.size++
}

func (hm *HashMap) Get(key interface{}) interface{} {
	index := hm.hash(fmt.Sprintf("%v", key))
	node := hm.table[index]
	for node != nil {
		if node.Key == key {
			return node.Value
		}
		node = node.Next
	}
	return nil
}
