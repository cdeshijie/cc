/* package main

import "fmt"

func main() { //defer延迟执行，在程序执行完在按照defer顺序逆序执行，可以看做是遇到defer就进栈，之后出栈执行
	println("start")
	defer println("step1")
	defer println("step2")
	defer println("step3")
	println("end")

	//接下来测试指针,指向数组的指针
	a := [3]int{1, 2, 3}
	//b:=&a
	//var b *[3]int
	//b = &a
	b := &a
	fmt.Printf("b: %v\n", b)

	//指针数组
	var c [3]*int
	for i, _ := range a {
		c[i] = &a[i]
	}
}
*/