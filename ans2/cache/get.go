package cache

import "fmt"

func (c *Cache) Get(key int64) {
	fmt.Println()
	if _, ok := c.items[key]; ok {
		fmt.Printf("Value for key %d : %d \n", key, c.items[key].value)
		return
	}
	fmt.Printf("Value not found for key  %d : -1 \n", key)
}
