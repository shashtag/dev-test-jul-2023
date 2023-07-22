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

	cache.Put(1, 311, 223)
	cache.Get(1)
	cache.Put(2, 4545, 6867)
	cache.Put(3, 42, 2)
	cache.Put(4, 24, 6)
	cache.Get(4)
	cache.Put(5, 2323, 1)
	cache.Get(3)
	cache.Get(1)
	cache.Put(6, 434, 5)

}
