package cache

import (
	"container/heap"
	"fmt"
	"math"
	"time"
)

func (c *Cache) Put(key, value, weight int64) {
	fmt.Println()
	current_time := time.Now().UnixNano()

	// if key exists, update the value and weight
	if _, ok := c.items[key]; ok {
		c.items[key] = cacheData{weight: weight, value: value, last_modified: current_time}
		fmt.Printf("Item %d updated \n", key)
		return
	}

	// if cache is full, pop the least score item
	if len(c.items) == c.capacity {

		for i := 0; i < len(heapDataList); i++ {
			cacheItem := c.items[heapDataList[i].key]

			heapDataList[i].score = float64(cacheItem.weight) / (math.Log(float64(current_time-cacheItem.last_modified+1)) + 1.00)
		}

		heap.Init(&heapDataList)
		poppedKey := heap.Pop(&heapDataList)
		delete(c.items, poppedKey.(int64))
		fmt.Printf("Item %d popped out of the chache \n", poppedKey)

	}

	// add the new item into heap list and cache map
	heapDataList = append(heapDataList, heapData{key: key})
	c.items[key] = cacheData{weight: weight, value: value, last_modified: current_time}

	fmt.Printf("Item %d added to the cache \n", key)

}
