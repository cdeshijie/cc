/* package main

import "fmt"

func main() {

	//map声明本质只有两种，一种make，一种不用make，不用make可以初始化赋值，这里四种区别是用没用var

	//第一种声明方式  []中是索引的类型，[]之后是数据的类型
	/* var m1 map[string]string     //此时m1只是空指针
	m1 = make(map[string]string) //make后申请了空间

	//第二种声明方式
	m2 := make(map[int]int) //第一种缩写

	//带容量的声明
	m3 := make(map[string]string, 10)

	//第三种声明方式,带初始化的声明，m1，m2，m3都是空集合，m4不是
	var m4 = map[string]int{"abc": 3, "ccd": 4}

	//第四种声明方式，make申请空间无法赋值，第三种缩写
	m5 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	//增
	m4["aaa"] = 5
	//删
	delete(m5, "orange")
	//改
	m4["abc"] = 1
	//查
	v2, ok := m5["aaa"] //v2的值是索引"aaa"的值，ok是bool型，有就是true，没有就是false，当没有时，v2就是各类型的0值
	fmt.Printf("v2: %v\n", v2)
	fmt.Printf("ok: %v\n", ok)

	//遍历
	for k, v := range m5 {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}*/
 