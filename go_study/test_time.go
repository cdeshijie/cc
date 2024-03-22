/*package main

import (
	"fmt"
	"time"
)

func test1() {
	t := time.Now() //time.Time是个结构体
	fmt.Printf("t: %T\n", t)
	fmt.Printf("t: %v\n", t)
	fmt.Printf("t.Year(): %v\n", t.Year())
}

func test2() { //测试时间戳
	now := time.Now()
	fmt.Printf("now.Unix(): %v\n", now.Unix())
	fmt.Println("-----------")
	timestamp := time.Now().Unix()     //时间戳转为时间
	timeobj := time.Unix(timestamp, 0) //转为时间
	fmt.Printf("timeobj: %v\n", timeobj)
}

func test3() { //测试时间的添加删除
	time1 := time.Now()
	time2 := time1.Add(time.Hour)
	fmt.Println(time2.Sub(time1))
}

func main() {
	test2()
}*/
