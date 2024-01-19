# Generic Heap Package

A flexible and customizable generic heap package for Go that empowers you to create heaps
with any type of values, providing the freedom to customize the heap structure based on
your specific requirements.


## Installation

```sh
go get github.com/SurajKadam7/heap
```

## Example : 
```go 
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

	minHeap := heap.New[Rank](minHeapFunc, 0)

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
```