package heaps

type Heap struct {
	Items []int
}

// -----------------------------------------------------------------------------

func NewHeap() Heap {
	return Heap{
		Items: make([]int, 0),
	}
}

// -----------------------------------------------------------------------------

func (heap *Heap) Insert(in int) {
	heap.Items = append(heap.Items, in)
	position := len(heap.Items)


	for {
		if position == 1 {
			break
		}
		parentPosition := position >> 1
		if heap.Items[position-1] >= heap.Items[parentPosition-1] {
			break
		}
		heap.Items[position-1], heap.Items[parentPosition-1] = heap.Items[parentPosition-1], heap.Items[position-1]
		position = parentPosition
	}
}

func (heap *Heap) ExtractMin() (out int, ok bool) {
	if len(heap.Items) == 0 {
		return out, false
	}
	out = heap.Items[0]
	if len(heap.Items) == 1 {
		heap.Items = heap.Items[:0]
		return out, true
	}
	heap.Items[0] = heap.Items[len(heap.Items)-1]
	heap.Items = heap.Items[:len(heap.Items)-1]

	if len(heap.Items) == 1 {
		return out, true
	}

	position := 1
	child1 := 2
	child2 := 3
	if len(heap.Items) == 2 {
		child2 = child1
	}

	for {
		if (heap.Items[position-1] <= heap.Items[child1-1]) && heap.Items[position-1] <= heap.Items[child2-1] {
			break
		}

		if heap.Items[child1-1] < heap.Items[child2-1] {
			heap.Items[position-1], heap.Items[child1-1] = heap.Items[child1-1], heap.Items[position-1]
			position = child1
		} else {
			heap.Items[position-1], heap.Items[child2-1] = heap.Items[child2-1], heap.Items[position-1]
			position = child2
		}

		child1 = position << 1
		if child1 > len(heap.Items) {
			break
		}
		child2 = position<<1 + 1
		if child2 > len(heap.Items) {
			break
		}
	}

	return out, true
}
