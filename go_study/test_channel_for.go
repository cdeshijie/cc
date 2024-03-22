/*package main

import "fmt"

var c = make(chan int) //无缓冲channel，此时只能读一个写一个交替进行,channel有一个写端，一个读端，只有两个都存在，channel空了，才能写入下一个数据
//假如只有读端，没有写端此时channel为空，读端阻塞；只有写段没有读端，因为是无缓冲channel，所以没法在写，写段阻塞，等待读端拿走数据后写段才能阻塞解除然后在写入,具体可以看网页收藏的channel有无缓冲区别
func main() {
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Printf("写入i: %v\n", i)
		}
		//close(c) //没有关闭时，前两次正常读写，第三次报错死锁，在关闭后第三次再从c取值就是0
	}()
	var k int
	for i := 0; i < 5; i++ { //这是知道channel里元素数量的时候的遍历
		k = <-c
		fmt.Printf("读出k: %v\n", k)
	}

	for { //这是不知道有多少元素时进行的遍历
		v, ok := <-c
		if ok {
			fmt.Printf("v: %v\n", v)
		} else {
			break
		}
	}
	for v := range c { //for range也可以遍历channel
		fmt.Printf("v: %v\n", v)
	}
}*/
