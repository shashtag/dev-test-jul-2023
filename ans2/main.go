package main

import (
	"fmt"
	"time"

	"github.com/shashtag/dev-test-jul-2023.git/ans2/cache"
)

func main() {
	cache := cache.NewCache(4)

	fmt.Println(cache)
	fmt.Println(time.Now().UnixNano())

	cache.Put(1, 1, 1)
	cache.Put(2, 1, 1)
	cache.Put(3, 1, 1)
	cache.Put(4, 1, 1)
	cache.Put(5, 1, 1)
	cache.Put(6, 1, 1)
}
