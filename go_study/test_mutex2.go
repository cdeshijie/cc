/*package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func main() { //不加锁前这里不是十万的原因我的理解是：假如某时刻i=50000，这个时刻同时有两个线程进行加操作，都拿的是50000，计算出都是50001，所以相当于两次加操作赋值后却是50001
	var count = 0
	var wg sync.WaitGroup
	//十个协程数量
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			//1万叠加
			for j := 0; j < 10000; j++ {
				lock.Lock()
				count++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}*/
