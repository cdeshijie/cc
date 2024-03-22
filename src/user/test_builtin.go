/*package main

import "fmt"

//make和new的不同:1.make只能用来分配和初始化类型为slice，map，channel的数据类型，new所有类型都可以
//2.new返回的是指针，make返回数据类型本身
//3.new分配的空间会被清零，make会初始化被分配的空间
func testnew1() { //测试new基本数据类型
	a := new(bool) //a,b,c皆为指针
	b := new(int)
	c := new(string)
	fmt.Printf("a: %v\n", *a) //此时ab皆为0值,c为空字符串
	fmt.Printf("b: %v\n", *b)
	fmt.Printf("c: %v\n", c)
}

func testnew2() { //测试new引用数据类型
	s := new([]int) //new出来的引用型变量不是nil，是一个空的数据类型的指针，应该是和make([]int,0)差不多,new不能提供初始化，make可以，make返回并非指针
	ss := make([]int, 0)
	fmt.Printf("ss: %v\n", ss) //ss: []，这里可以看出，从内容上讲，*s和ss没有太大区别，都是一个空的切片，就是一个是指针，一个是数据类型本身
	if s == nil {
		fmt.Println("nil")
	}
	fmt.Printf("s: %T\n", s)             //s: *[]int
	fmt.Printf("s: %v\n", *s)            //s: []
	fmt.Printf("len(*s): %v\n", len(*s)) //len(*s): 0
	*s = append(*s, 1)                   //这里要用*s，即切片本身
	fmt.Printf("s: %v\n", *s)            //s: [1]
	*s = make([]int, 10)                 //make后，原来里面的东西就没有了
	*s = append(*s, 3)
	fmt.Printf("s: %v\n", *s) //s: [0 0 0 0 0 0 0 0 0 0 3]

	var s1 []int //此时s1是nil
	if s1 == nil {
		fmt.Println("nil")
	}
}

func testpanic() {
	defer fmt.Println("panic前会先执行defer")
	panic("error!!!!")
	fmt.Println("我不会执行")
}

func main() {
	testnew2()
}
*/