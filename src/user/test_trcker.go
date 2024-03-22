/* package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second) //timer只能用一次，ticker可以循环使用
	/*for _ = range t.C {              //每隔一秒，输出一次
		fmt.Println("ticker.....")
	}

	c := make(chan int)
	go func() {
		for _ = range t.C {
			select {
			case c <- 1:
				fmt.Println("c收到：1")
			case c <- 2:
				fmt.Println("c收到：2")
			case c <- 3:
				fmt.Println("c收到：3")
			}
		}
	}()
	sum := 0
	for v := range c { //正常channelrange前需要关闭通道，这里没关，但是因为是无缓存channel所以没事。forrange会取出channel里的数据，而不是拷贝
		fmt.Printf("v: %v\n", v)
		sum += v
		fmt.Printf("len(c): %v\n", len(c))
		fmt.Printf("sum: %v\n", sum)
		if sum >= 10 {
			break
		}
	}
	//fmt.Printf("sum: %v\n", sum)
}
*/