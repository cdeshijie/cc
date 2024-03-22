/* package main

import "fmt"

type USBER interface { //接口命名一般以er结尾
	write()
	read()
}
type computer struct {
	name string
}

func (c computer) write() { //write是接口里的方法，这里重写write并归属于computer，所以是属于调用了接口USBER，在调用接口时，必须把接口的所有方法都实现，不能只实现write或者read
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("write...")
}
func (c computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("read....")
}
func main() {
	c := computer{
		"lianxiang",
	}
	c.read()
	c.write()
}
*/