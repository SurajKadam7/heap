package heap

// heap formula
// left child :- 2*i (where i is the parent index)
// right child :- 2*i+1
// parent :- n/2 (where n is the child node index)

// HeapFunc will table three arguments and return bool as a result
// exmple of HeapFunc
// MinHeap  := func(heap []int, parent, child int) bool { return heap[child] >= heap[parent] }
// MaxHeap  := func(heap []int, parent, child int) bool { return heap[child] <= heap[parent] }
type HeapFunc[T any] func(heap []T, parent, child int) bool

type heap[T any] struct {
	data     []T
	heapFunc HeapFunc[T]
}

func New[T any](f HeapFunc[T], size int) *heap[T] {
	size += 1
	return &heap[T]{
		data:     make([]T, size),
		heapFunc: f,
	}
}

func (heap *heap[T]) Push(val T) {
	h := heap.data
	h = append(h, val)
	child := len(h) - 1
	parent := child / 2

	for parent > 0 && !heap.heapFunc(h, parent, child) {
		h[parent], h[child] = h[child], h[parent]
		child = parent
		parent = parent / 2
	}
	heap.data = h
}

func (h *heap[T]) Pop() (value T) {
	if len(h.data) <= 1 {
		return
	}

	lst := len(h.data) - 1 // len = elements + 1
	parent := 1

	value = h.data[parent]
	h.data[parent], h.data[lst] = h.data[lst], h.data[parent]
	h.data = h.data[:lst]

	h.balance(parent)
	return value
}

func Heapify[T any](arr []T, f HeapFunc[T]) *heap[T] {
	h := New[T](f, len(arr))

	for parent := len(arr); parent > 0; parent-- {
		h.data[parent] = arr[parent-1]
		h.balance(parent)
	}

	return h
}

func (heap *heap[T]) balance(parent int) {
	h := heap.data

	for {
		left, right, balanceParent := parent*2, parent*2+1, parent

		leftSwap := false
		if left < len(h) && !heap.heapFunc(h, parent, left) {
			h[left], h[parent] = h[parent], h[left]
			leftSwap = true
			balanceParent = left
		}

		if right < len(h) && !heap.heapFunc(h, parent, right) {
			h[right], h[parent] = h[parent], h[right]
			if leftSwap {
				h[left], h[right] = h[right], h[left]
			}
			balanceParent = right
		}

		if balanceParent == parent {
			break
		}

		parent = balanceParent
	}
}
