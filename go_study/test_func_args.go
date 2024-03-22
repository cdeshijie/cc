/* package main

import "fmt"

func f1(a []int) { //这里和C语言不同，参数数组长度未说明就是切片，只能传入切片
	a[0] = 100
}

func f2(a [2]int) { //因为传入是数组，不是引用类型，所以不改变参数里的值
	a[0] = 100
}

func f3(a ...int) { //参数是任意长度的int型变量
	for _, v := range a {
		fmt.Printf("v: %v\n", v)
	}
}
func main() {
	a := [2]int{1, 2} //a是数组，可以f2,但是不能f1
	f2(a)
	fmt.Printf("a: %v\n", a) //此时a值不变

	b := []int{1, 2, 3} //b是切片
	f1(b)
	fmt.Printf("b: %v\n", b) //此时b里的参数发生了改变

	//f3(a) 此时报错，说明不能用固定长度的数组来当 参数的长度为非固定的函数 的参数
	f3(a[0], a[1])
	f3(b...)
}
*/