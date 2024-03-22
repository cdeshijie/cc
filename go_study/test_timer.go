/* package main

import (
	"fmt"
	"time"
)

//var x time.Timer=*time.NewTimer(time.Second*2) 主函数外声明

func main() { //延迟执行的n种方法
	//延迟两秒
	t := time.NewTimer(time.Second * 2) //t是一个timer计时器，意思是两秒后在timer里的c写入当前时间
	fmt.Printf("time.Now(): %v\n", time.Now())
	t1 := <-t.C //timer.c是一个channel，如上所述，在调用时会等待设置的时间并到了时间后写入当前时间进c，这里t1取出来
	fmt.Printf("t1: %v\n", t1)

	//延迟两秒
	fmt.Printf("time.Now(): %v\n", time.Now())
	t2 := time.NewTimer(time.Second * 2)
	<-t2.C //只有执行c操作才会延时
	fmt.Printf("time.Now(): %v\n", time.Now())

	//延迟两秒
	fmt.Printf("time.Now(): %v\n", time.Now())
	<-time.After(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())

	//延迟两秒
	fmt.Printf("time.Now(): %v\n", time.Now())
	time.Sleep(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())

	//这段测试timer.stop作用，
	t3 := time.NewTimer(time.Second) //创建计时器
	go func() {
		<-t3.C                  //延迟1s，说白了就是阻塞1s
		fmt.Println("func....") //阻塞完输出
	}()
	t3.Stop()                   //t3停止计时，即永远不会归0，所以func不会被打印了
	time.Sleep(time.Second * 3) //main等待3s，否则main执行完func直接被结束
	fmt.Println("end....")      //main执行结束

	t4 := time.NewTimer(time.Second * 2) //测试reset，阻塞2s
	fmt.Printf("time.Now(): %v\n", time.Now())
	t4.Reset(time.Second * 4) //重新设置，阻塞4s
	<-t4.C
	fmt.Printf("time.Now(): %v\n", time.Now())
}
*/