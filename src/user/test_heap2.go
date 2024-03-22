package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 首先明确的是heap包里的堆的插入和删除方法是写好的了，但是需要使用传入参数的less，len，swap，pop，push五个方法，参数的pop和heap.pop不是一回事，不懂去看源码
// 要想使用heap包里的堆的插入和删除方法，则需要有五个方法的实现，less，len，swap，pop，push，sort.IntSlice提供了less，len，swap，所以自己需要实现pop，push
var factors = []int{2, 3, 5}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	/*a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	fmt.Println(h.IntSlice)
	return v*/
	v := h.IntSlice[len(h.IntSlice)-1]
	//h.IntSlice = append(h.IntSlice, h.IntSlice[:len(h.IntSlice)-1]...),这样写是错的,h.IntSlice=append(h.IntSlice[0:0],h.IntSlice[:len(h.IntSlice)-1]...)这样写就对了
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	fmt.Println(h.IntSlice)
	return v
}

func nthUglyNumber(n int) int {
	h := &hp{sort.IntSlice{1}}
	seen := map[int]struct{}{1: {}}
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}
func main() {
	//fmt.Println(nthUglyNumber(5))
	a := []int{1, 2, 3}
	a = append(a, a[:len(a)-1]...)
	fmt.Println(a)
}
