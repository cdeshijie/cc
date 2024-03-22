/*package main

import "fmt"

var chanint = make(chan int)
var chanstring = make(chan string)

func main() {
	go func() {
		defer close(chanint)
		defer close(chanstring)
		chanint <- 10
		chanstring <- "hello"
	}()
	for { //无限循环
		select { //类似Switch，这里是控制channel的，两个case哪个满足执行哪个，都满足随机挑选执行，假如没有close，两个case都不满足执行default，有close则随机执行取0或者取nil输出(取默认值)
		case r := <-chanint:
			fmt.Printf("r: %v\n", r)
		case r := <-chanstring:
			fmt.Printf("r: %v\n", r)
		default:
			fmt.Println("default.....")
		}
	}
}*/
