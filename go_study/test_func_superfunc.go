/*package main

import "fmt"

func sayhello(name string) {
	fmt.Printf("hello, %v\n", name)
}

func test1(name string, f func(string)) { //函数作为参数时，如果不声明返回值就是无返回值函数，即看实际使用，不能不声明但是有返回值
	f(name)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(c string) func(int, int) int { //函数在作为另一个函数里return返回值时，可以声明其返回值类型,也可以不声明，不声明就是无返回值函数，声明就是有返回值函数，不能不声明返回值但是有返回值
	switch c {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func test2(c string, x, y int) {
	//fmt.Printf("cal(c)(x, y): %v\n", cal(c)(x, y))  //cal(c)(x,y)是下面两句的集合
	var t = cal(c)
	fmt.Printf("t(x, y): %v\n", t(x, y))
}

func main() {
	//test1("TOM", sayhello)
	test2("+", 1, 2)
}*/
