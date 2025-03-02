type MaxHeap []int

func (h MaxHeap) Len() int {
    return len(h)
}
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h MaxHeap) Less(i, j int) bool {
    return h[i] > h[j] 
} // Max heap
func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0: n-1]
    return x
}

type MinHeap []int

func (h MinHeap) Len() int {
    return len(h)
}
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0: n-1]
    return x
}

type MedianFinder struct {
    maxHeap *MaxHeap
    minHeap *MinHeap
}


func Constructor() MedianFinder {
    maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	
	return MedianFinder{
		maxHeap: maxHeap,
		minHeap: minHeap,
	}
}


func (this *MedianFinder) AddNum(num int)  {
    if this.maxHeap.Len() == 0 || num <= (*this.maxHeap)[0] {
        heap.Push(this.maxHeap, num)
    } else {
        heap.Push(this.minHeap, num)
    }

    if this.maxHeap.Len() > this.minHeap.Len()+1 {
        heap.Push(this.minHeap, heap.Pop(this.maxHeap))
    } else if this.minHeap.Len() > this.maxHeap.Len() {
        heap.Push(this.maxHeap, heap.Pop(this.minHeap))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Len() > this.minHeap.Len() {
        return float64((*this.maxHeap)[0])
    }
    return float64((*this.maxHeap)[0] + (*this.minHeap)[0]) / 2.0
}