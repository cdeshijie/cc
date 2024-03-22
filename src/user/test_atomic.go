/* package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var i int = 100
var lock sync.Mutex //加锁
var k int32 = 100

func add() {
	i++
}

func sub() {
	i--
}

func add1() { //加锁后，i++只能同时由一个协程进行使用，所以每次都不会出现同时取出相同值的问题
	lock.Lock()
	i++
	lock.Unlock()
}

func sub1() {
	lock.Lock()
	i--
	lock.Unlock()
}

func add2() { //atomic为原子操作，不可分割，确保某个变量在同一时间只有一个协程去操作，避免大量锁的使用
	atomic.AddInt32(&k, 1)
}

func sub2() {
	atomic.AddInt32(&k, -1)
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second * 2) //2s后协程都执行结束
	fmt.Printf("i: %v\n", i)    //不加锁时，按理说add和sub都执行100次，但是最后值不是100，因为在某时刻i假设为100，同时进行两次sub操作，取得都是100-1=99，再把99赋予i，相当于两次sub却只减了1.
	i = 100
	for i := 0; i < 100; i++ {
		go add1()
		go sub1()
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("i: %v\n", i)

	for i := 0; i < 100; i++ {
		go add2()
		go sub2()
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("k: %v\n", k)
}
*/