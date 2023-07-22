package cache

type cacheData struct {
	weight        int64
	last_modified int64
	value         int64
}

type Cache struct {
	capacity int
	items    map[int64]cacheData
}

func NewCache(capacity int) *Cache {
	heapDataList = make([]heapData, 0, capacity)
	return &Cache{capacity: capacity, items: make(map[int64]cacheData)}
}
