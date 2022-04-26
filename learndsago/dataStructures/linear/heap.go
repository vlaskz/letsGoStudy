package linear

import (
	"container/heap"
	"fmt"
)

//this guy represents an array of ints.

type IntegerHeap []int

func (iheap IntegerHeap) Len() int { return len(iheap) }

func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int

	var prev IntegerHeap = *iheap
	n = len(prev)
	x1 = prev[n-1]
	*iheap = prev[0 : n-1]
	return x1

}

func heap_example() {
	var i *IntegerHeap = &IntegerHeap{1, 3, 2, 5, 4, 8, 7}

	heap.Init(i)
	heap.Push(i, 99)
	fmt.Printf("minimum: %d\n", (*i)[0])

	for i.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(i))
	}
}
