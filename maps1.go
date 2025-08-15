package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]string
	mu   sync.Mutex
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func main() {
	cache := &Cache{
		data: make(map[string]string),
	}

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	value, ok := cache.Get("key1")
	if ok {
		fmt.Println("Value for key1:", value)
	} else {
		fmt.Println("Key1 not found in cache")
	}

	value, ok = cache.Get("key2")
	if ok {
		fmt.Println("Value for key2:", value)
	} else {
		fmt.Println("Key2 not found in cache")
	}

	value, ok = cache.Get("key3")
	if ok {
		fmt.Println("Value for key3:", value)
	} else {
		fmt.Println("Key3 not found in cache")
	}
}