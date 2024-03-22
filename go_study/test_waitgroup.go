/*package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showMsg(i int) {
	defer wp.Done() //WaitGroup计数-1.
	fmt.Printf("i: %v\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wp.Add(1) //每开一个协程，WaitGroup计数+1。
		go showMsg(i)
	} //因为先执行for循环，所以当计数器+1后才会开线程，线程开启后for又会循环，此时上个线程才刚开启，这里第二轮for循环又开始了，就保证WaitGroup计数器在所有协程执行完之前不会提前为0
	wp.Wait() //等待WaitGroup计数为0
	fmt.Println("end....")
}*/
