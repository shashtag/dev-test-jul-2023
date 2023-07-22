package cache

type heapData struct {
	key   int64
	score float64
}

type HeapDataHeap []heapData

var heapDataList HeapDataHeap

func (h HeapDataHeap) Len() int           { return len(h) }
func (h HeapDataHeap) Less(i, j int) bool { return h[i].score < h[j].score }
func (h HeapDataHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HeapDataHeap) Push(x any) {
	*h = append(*h, x.(heapData))
}

func (h *HeapDataHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x.key
}
