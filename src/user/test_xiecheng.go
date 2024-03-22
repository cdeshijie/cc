/*package main

import (
	"fmt"
	"strconv"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

// 编写一个程序完成如下功能：1.主线程开启一个协程，每秒输出一次“hello golang”
// 主线程每秒输出一次“hello msb”，10s后结束
// 要求主线程和协程同时进行
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() { //函数添加go关键字表示新建一个协程来执行该函数，没有加go在main函数里表示在main函数协程了，这里两个go一个没有，表示有三个协程
	//go showMsg("java")
	//go showMsg("python")
	//fmt.Println("main...")
	//time.Sleep(time.Millisecond * 2000)
	go test()//go 不能放for后面，因为go是协程，for执行完主线程就执行完了，次数协程才刚开始，但是主线程结束就是结束协程，所以协程也不会执行
	for i := 0; i < 10; i++ {
		fmt.Println("hello msb" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}*/
