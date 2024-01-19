package main

import (
	"fmt"
)

func main() {
	// heap := []int{}
	// heap = append(heap, 0)
	// heap = Push(heap, 50)
	// heap = Push(heap, 30)
	// heap = Push(heap, 20)
	// heap = Push(heap, 15)
	// heap = Push(heap, 10)
	// heap = Push(heap, 8)
	// heap = Push(heap, 16)
	// fmt.Println(heap)
	// var val int
	// for len(heap) > 1 {
	// 	val, heap = Pop(heap)
	// 	fmt.Println(val)
	// }
	// fmt.Println(heap)

	h := New(MinHeap)
	h.Push(50)
	h.Push(30)
	h.Push(20)
	h.Push(15)
	h.Push(10)
	h.Push(8)
	h.Push(16)

	for len(h.data) > 1 {
		fmt.Println(h.Pop())
	}
	fmt.Println(h.data)

}
