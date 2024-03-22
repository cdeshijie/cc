/* package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

var v = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10) //获取从0到10的随机数
	fmt.Printf("send:%d\n", value)
	//time.Sleep(time.Second * 5)
	values <- value //随机数放入通道
}

func main() { //
	defer close(values) //最后关闭通道
	go send()           //往通道送数据
	fmt.Println("wait....")
	value := <-values                //从通道取数据,send协程中给通道送入数据后，main里才能得到，所以会一直等待，不会完成main后直接结束所有协程
	fmt.Printf("recive:%d\n", value) //输出得到的数据是多少
	fmt.Println("end.....")
}
*/