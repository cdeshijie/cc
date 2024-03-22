/* package main

import "fmt"

func sum1(a int, b int) (ret int) { //ret为返回值,带参数带返回值的函数
	ret = a + b
	return ret
}

func sum2(a int, b int) int { //带参数带返回值，这里只声明返回值的类型
	ret := a + b
	return ret //return a+b
}

func f1() (name string, age int) { //带参数，没有返回值
	name = "tom"
	age = 16
	return name, age //返回值必须按照声明的顺序
}

func f2() (name string, age int) { //return的另一种使用
	name = "tom"
	age = 16
	return
}
func f3() (name string, age int) { //返回值不用声明好的，而是用函数里自己声明的，此时只要类型相同就可以
	a := "tom"
	b := 16
	return a, b
}
func main() {
	name, age := f2()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

	name1, _ := f3() //此时只要一个返回值
	fmt.Printf("age2: %v\n", name1)
}
*/