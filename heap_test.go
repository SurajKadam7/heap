package heap

import (
	"fmt"
	"math/rand"
	"testing"
)

type Comparables interface {
	int | int8 | int32 | int64 | uint | uint8 | uint32 | uint64 | float32 | float64
}

func getRandomInt[T Comparables]() T {
	return T(rand.Intn(9000000) + 1)
}

var (
	MinHeap HeapFunc[int] = func(heap []int, parent, child int) bool { return heap[parent] <= heap[child] }
	MaxHeap HeapFunc[int] = func(heap []int, parent, child int) bool { return heap[parent] >= heap[child] }
)

func test[T Comparables](h *heap[T], f HeapFunc[T], heapType string, size int) error {
	for i := 1; i < size+1; i++ {
		parent := i
		left, right := 2*parent, 2*parent+1
		if left < len(h.data) && !f(h.data, parent, left) {
			return fmt.Errorf("does not follow the %s heap rule parent : %v left : %v", heapType, h.data[parent], h.data[left])
		}

		if right < len(h.data) && !f(h.data, parent, right) {
			return fmt.Errorf("does not follow the %s heap rule parent : %v right : %v", heapType, h.data[parent], h.data[right])
		}
	}
	return nil
}

func getMinHeap[T int](size int) (*heap[T], HeapFunc[T]) {
	MinHeap := func(heap []T, parent, child int) bool { return heap[parent] <= heap[child] }
	hMin := New[T](MinHeap)
	for i := 0; i < size; i++ {
		hMin.Push(getRandomInt[T]())
	}
	return hMin, MinHeap
}

func getMaxHeap[T int](size int) (*heap[T], HeapFunc[T]) {
	MaxHeap := func(heap []T, parent, child int) bool { return heap[parent] >= heap[child] }
	hMax := New[T](MaxHeap)
	for i := 0; i < size; i++ {
		hMax.Push(getRandomInt[T]())
	}
	return hMax, MaxHeap
}

//-----------------------------------------------------------------------------------

func Test_heap_Push(t *testing.T) {
	sz := 2000000
	hMin, MinHeap := getMinHeap[int](sz)
	if err := test[int](hMin, MinHeap, "minHeap", sz); err != nil {
		t.Error(err)
	}

	hMax, MaxHeap := getMaxHeap[int](200)
	if err := test[int](hMax, MaxHeap, "maxHeap", sz); err != nil {
		t.Error(err)
	}
}

func Test_heap_Pop(t *testing.T) {
	sz := 2000000
	var second int

	hMin, _ := getMinHeap[int](sz)
	first := hMin.Pop()

	for i := 0; i < sz-1; i++ {
		second = hMin.Pop()
		if second < first {
			break
		}
		first = second
	}

	if len(hMin.data) > 1 && second != 0 {
		t.Errorf("min heap does not satisfy firstPop: %v secondPop: %v", first, second)
	}

	hMax, _ := getMaxHeap[int](sz)
	first = hMax.Pop()
	for i := 0; i < sz-1; i++ {
		second = hMin.Pop()
		if second > first {
			break
		}
		first = second
	}

	if len(hMin.data) > 1 && second != 0 {
		t.Errorf("max heap does not satisfy firstPop: %v secondPop: %v", first, second)
	}
}

func TestHeapify(t *testing.T) {
	sz := 2000000
	arr := make([]int, sz)
	for i := 0; i < len(arr); i++ {
		arr[i] = getRandomInt[int]()
	}

	MinHeap := func(heap []int, parent, child int) bool { return heap[parent] <= heap[child] }
	hMin := Heapify[int](arr, MinHeap)
	if err := test[int](hMin, MinHeap, "minHeap", sz); err != nil {
		t.Error(err)
	}

	MaxHeap := func(heap []int, parent, child int) bool { return heap[parent] >= heap[child] }
	hMax := Heapify[int](arr, MaxHeap)
	if err := test[int](hMax, MaxHeap, "maxHeap", sz); err != nil {
		t.Error(err)
	}
}
