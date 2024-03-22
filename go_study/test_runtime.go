/* package main

import (
	"fmt"
	"runtime"
	"time"
)

func showMsg(s string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("s: %v\n", s)
	}
}

func showMsg2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			runtime.Goexit() //如果i大于等于5，协程直接结束并退出
		}
	}
}

func a() {
	for i := 0; i < 15; i++ {
		fmt.Println("a:", i)
	}
}

func b() {
	for i := 0; i < 15; i++ {
		fmt.Println("b:", i)
	}
}

func main() {
	/*go showMsg("java")//测试runtime.Gosched
	runtime.Gosched() //gosched的作用是在轮到main协程执行时把cpu让给其他协程，所以现在必然等showmas完成后才会输出end....
*/

/*go showMsg2()测试runtime.Goexit
time.Sleep(time.Second) //加了main睡眠后，一秒钟足以让协程showmag2执行完成,在输出end....
fmt.Println("end.....")*/

/*fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU()) //输出cpu个数，现在是4，所以在运行协程时，哪个空闲用哪个,测试runtime.GOMAXPROCS
	go a()
	go b()                //没有使用runtime.GOMAXPROCS前，a和b可能a先执行，也可能b先执行，也可能ab交替执行
	runtime.GOMAXPROCS(1) //最大cpu核心设为1,这样的话在不会出现ab交替执行的情况，可能先执行a，也可能先执行b，但是是在执行完某一个后再去执行另一个
	go a()
	go b()
	time.Sleep(time.Second * 2)
	fmt.Println("end...")
}
*/