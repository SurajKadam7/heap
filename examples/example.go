package main

import (
	"fmt"

	"github.com/SurajKadam7/heap"
)

type Rank struct {
	Name  string
	Score int
}

func main() {

	minHeapFunc := func(heap []Rank, parent, child int) bool {
		return heap[parent].Score < heap[child].Score
	}

	minHeap := heap.New[Rank](minHeapFunc)

	minHeap.Push(Rank{
		Name:  "ABC",
		Score: 30,
	})

	minHeap.Push(Rank{
		Name: "EFG",
		Score: 20,
	})

	fmt.Println(minHeap.Pop())
	fmt.Println(minHeap.Pop())
	fmt.Println(minHeap.Pop())
}
