/*package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100 //函数外只能用var，不能用：=
var n int = 0
var wp sync.WaitGroup

func Add() {
	i += 1
	fmt.Printf("i++: %v\n", i)
}

func Sub() {
	i -= 1
	fmt.Printf("i--: %v\n", i)
}

func Add1() {
	defer wp.Done()
	i += 1
	fmt.Printf("i++: %v\n", i)
	n++
}

func Sub1() {
	defer wp.Done()
	i -= 1
	fmt.Printf("i--: %v\n", i)
	n++
}

func main() {
	/*for i := 0; i < 50; i++ { //此时会出现，只加了一次但是i等于102的情况，因为开了新协程后，可能会出现两个协程加了但是第一个协程打印了i的值，此时为102，出现错误
		go Add()
		go Sub()
	}
	fmt.Println("endi:", i) //并且还会出现endi等于101的情况，在输出endi的时候，有协程没有结束，自然endi不会是100*/

/*for k := 0; k < 300; k++ { //此时endi出现了101的情况，说明中间
		wp.Add(1)
		go Add1()
		wp.Add(1)
		go Sub1()
	}
	wp.Wait() //正常来说main协程会阻塞，wait会等wp的所有协程执行完，咱们理解的是加减各50次，之后i一定是100
	//但是协程多了之后可能并非是所有协程执行完才是0，而是中间某一时刻正好为0，就出现了wait后但是endi不为100的情况
	fmt.Println("endi", i)
	fmt.Printf("n: %v\n", n)

	time.Sleep(time.Second) //这里让main函数沉睡以此完成剩下的进程，但是两次endi输出相同证明了在
	fmt.Println("endi", i)
}*/
