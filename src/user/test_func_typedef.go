/* package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func comp(a int, b int) int {
	if a > b {
		return a
	} else if a < b {
		return b
	} else {
		return a
	}

}

func test1(a int, y int) string {
	return "cc"
}

func main() {
	type fun func(int, int) int //定义一个类型为fun的变量，fun是一个函数，函数的参数为两个int，返回值需要声明清楚,定义新类型
	var f fun
	f = sum
	s := f(1, 2)
	fmt.Printf("s: %v\n", s)
	f = comp
	c := f(2, 3)
	fmt.Printf("c: %v\n", c)
	//f = test1  f是两个参数为int返回值也是int的变量，所以test1报错


	type INT = int //类型别名,两种type方式区别在于，别名可以调用源类型的方法，新类型不行
}
*/