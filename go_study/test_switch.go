/* package main

import "fmt"

func main() {
	a := 100
	switch { //switch用法 switch后不解变量可以在case后加判断语句 go语言默认只执行第一个符号条件的case不同于c语言 想要接着执行需要fallthrough但是也只执行一个
	case a == 100:
		fmt.Println("100")

	case a < 200:
		fmt.Println("200")
		fallthrough
	default:
		fmt.Println("其他值")
	}
}
*/