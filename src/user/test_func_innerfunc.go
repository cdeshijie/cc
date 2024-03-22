/*package main

import "fmt"

func main() {

	max := func(x, y int) int { //在函数里声明一个匿名函数,与type不同的是，type声明的是一种类型的变量，变量可以赋值为不同的函数，比如sum和comp，匿名函数则是一个函数，声明后就固定了
		if x > y {
			return x
		} else {
			return y
		}
	} //max是一个函数
	p := max(4, 3)
	fmt.Printf("p: %v\n", p)

	r := func(x, y int) int { //另一种调用函数的方式,r是一个int值，而非函数
		return x + y
	}(1, 2)
	fmt.Printf("r: %v\n", r)

}*/
