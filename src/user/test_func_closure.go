/*package main

import "fmt"

func f1() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		fmt.Printf("%v", &x)
		return x
	}
}

func main() {
	//第一种声明方法
	var a func(y int) int //带括号与不带括号的区别
	a = f1()
	var b func() func(y int) int
	b = f1

	//第二种声明方法
	//a := f1() //f此时是func(y int) int型变量
	//b := f1   //f2此时是func() func(y int) int型变量

	fmt.Printf(" %v\n", a(1)) //此时为闭包,闭包就是在函数内调用了函数外的值，这里是a调用了外面的x，所以只要a存在，x就会随着a的使用而改变,即保留了x的指针
	fmt.Printf(" %v\n", a(1))

	fmt.Printf("b()(1): %v\n", b()(1)) //正常调用,这里每次都会var x，所以x的地址发生了改变
	fmt.Printf("b()(1): %v\n", b()(1))

	c := f1() //这里c重新申请一块地址，所以x地址也改变了，即不会顺着a的值传递,切记:=是var的缩写
	fmt.Printf(" %v\n", c(1))
	fmt.Printf(" %v\n", c(1))
}*/
