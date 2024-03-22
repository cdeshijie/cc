/* package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {//结构体和数组一样都是值传递，会进行副本复制
	tom := student{ //按键值初始化,要么全键值，要么全不键值，不能混合使用，结构体在作为方法的参数时，参数非指针就是值传递，不改变外面结构体，指针则会改变外面结构体的值
		name: "tom",
		age:  16,
	} //var tom student = student{name:"tom",age:16}

	fmt.Printf("tom: %v\n", tom)
	jick := student{ //无键值初始化
		"jick",
		12,
	}
	fmt.Printf("jick: %v\n", jick)
	lilei := student{ //部分初始化必须指定键值
		name: "lilei",
	}
	fmt.Printf("lilei: %v\n", lilei)

	var p *student //结构体指针
	p = &tom
	fmt.Printf("p: %p\n", p) //输出指针
	fmt.Printf("p: %v\n", *p)

	p = new(student) //通过new来创建一个结构体给p
	p.name = "jery"  //结构体指针和结构体本身都可以通过.来调用结构体里面的参数
}
*/