/*package main

import (
	"container/heap"
	"fmt"
)

// 如果不使用sort.intslice来使用堆，就需要先实现五个方法
type MyHeap []int

func (h MyHeap) Len() int {
	return len(h)
}

func (h MyHeap) Less(i, j int) bool { //h[i] < h[j]是小根堆，h[i] > h[j]是大根堆
	return h[i] < h[j]
}

func (h MyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MyHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x

	//下面这种写法也可以
	/*x := (*h)[h.Len()-1]
	*h = append(*h, (*h)[:h.Len()-1]...)
	return x*/
/*}

func main() {
	h := &MyHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println("Top:", (*h)[0]) // 输出：Top: 1
	fmt.Println("Len:", h.Len()) // 输出：Len: 4
	fmt.Println(*h)
	t := heap.Pop(h)
	fmt.Println(t)
}*/
