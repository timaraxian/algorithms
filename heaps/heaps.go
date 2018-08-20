package heaps

type Heap struct {
	items []int
}

// -----------------------------------------------------------------------------

func NewHeap() Heap {
	return Heap{
		items: make([]int, 0),
	}
}

// -----------------------------------------------------------------------------

func (heap *Heap) Insert(in int) {
	position := len(heap.items) + 1
	heap.items = append(heap.items, in)

	parentIndex := position >> 1

	for {
		if len(heap.items) == 1 {
			break
		}
		if heap.items[position-1] >= heap.items[parentIndex-1] {
			break
		}
		heap.items[position-1], heap.items[parentIndex-1] = heap.items[parentIndex-1], heap.items[position-1]
		position = parentIndex
		parentIndex = position >> 1
	}
}

func (heap *Heap) ExtractMin() (out int, ok bool) {
	if len(heap.items) == 0 {
		return out, false
	}
	out = heap.items[0]
	if len(heap.items) == 1 {
		heap.items = heap.items[:0]
		return out, true
	}
	heap.items[0] = heap.items[len(heap.items)-1]
	heap.items = heap.items[:len(heap.items)-1]

	if len(heap.items) == 1 {
		return out, true
	}

	position := 1
	child1 := 2
	child2 := 3
	if len(heap.items) == 2 {
		child2 = child1
	}

	for {
		if (heap.items[position-1] <= heap.items[child1-1]) && heap.items[position-1] <= heap.items[child2-1] {
			break
		}

		if heap.items[child1-1] < heap.items[child2-1] {
			heap.items[position-1], heap.items[child1-1] = heap.items[child1-1], heap.items[position-1]
			position = child1
		} else {
			heap.items[position-1], heap.items[child2-1] = heap.items[child2-1], heap.items[position-1]
			position = child2
		}

		child1 = position << 1
		if child1 > len(heap.items) {
			break
		}
		child2 = position<<1 + 1
		if child2 > len(heap.items) {
			break
		}
	}

	return out, true
}
