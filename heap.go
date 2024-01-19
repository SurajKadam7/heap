package main

// heap formula
// left child :- 2*i (where i is the parent index)
// right child :- 2*i+1
// parent :- n/2 (where n is the child node index)

type HeapFunc func(heap []int, child, parent int) bool

var (
	MinHeap HeapFunc = func(heap []int, child, parent int) bool { return heap[child] > heap[parent] }
	MaxHeap HeapFunc = func(heap []int, child, parent int) bool { return heap[child] < heap[parent] }
)

type heap struct {
	data     []int
	heapFunc HeapFunc
}

func New(f HeapFunc) *heap {
	return &heap{
		data:     make([]int, 1),
		heapFunc: f,
	}
}
func (heap *heap) Push(val int) {
	h := heap.data
	h = append(h, val)
	child := len(h) - 1
	parent := child / 2

	for parent > 0 && heap.heapFunc(h, parent, child) {
		h[parent], h[child] = h[child], h[parent]
		child = parent
		parent = parent / 2
	}
	heap.data = h
}

func (heap *heap) Pop() int {
	h := heap.data
	lst := len(h) - 1 // len > elements + 1
	parent := 1

	pop := h[parent]
	h[parent], h[lst] = h[lst], h[parent]

	for {
		left, right, oldParnt := parent*2, parent*2+1, parent

		if left < len(h)-1 && heap.heapFunc(h, parent, left) {
			h[left], h[parent] = h[parent], h[left]
			parent = left
		}

		if right < len(h)-1 && heap.heapFunc(h, parent, right) {
			h[right], h[parent] = h[parent], h[right]
			parent = right
		}

		if oldParnt == parent {
			break
		}
	}

	h = h[:lst]
	heap.data = h

	return pop
}

func Push(h []int, val int) []int {
	h = append(h, val)
	child := len(h) - 1
	parent := child / 2

	for parent > 0 && h[parent] < h[child] {
		h[parent], h[child] = h[child], h[parent]
		child = parent
		parent = parent / 2
	}

	return h
}

func Pop(h []int) (int, []int) {
	lst := len(h) - 1 // len > elements + 1
	parent := 1

	pop := h[parent]
	h[parent], h[lst] = h[lst], h[parent]

	for {
		left, right, oldParnt := parent*2, parent*2+1, parent

		if left < len(h)-1 && h[parent] < h[left] {
			h[left], h[parent] = h[parent], h[left]
			parent = left
		}

		if right < len(h)-1 && h[parent] < h[right] {
			h[right], h[parent] = h[parent], h[right]
			parent = right
		}

		if oldParnt == parent {
			break
		}
	}

	// poping the element
	h = h[:lst]
	return pop, h
}

// func (h heap) Heapify(arr []int) {}
// [0 50 30 20 8 15 10 16]
