package cache

import (
	"fmt"
	"time"
)

func (c *Cache) Get(key int64) {
	fmt.Println()

	// Check if key exists in cache or not and return value if exists also update last_modified time
	// Time - O ( 1 )
	if data, ok := c.items[key]; ok {
		fmt.Printf("Value for key %d : %d \n", key, data.value)
		c.items[key] = cacheData{weight: data.weight, value: data.value, last_modified: time.Now().UnixNano()}
		return
	}
	fmt.Printf("Value not found for key  %d : -1 \n", key)
}
