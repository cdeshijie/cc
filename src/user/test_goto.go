/* package main

import "fmt"

func main() {
lable:
	fmt.Print("1\n") //goto到的lable与continue和break的lable不同，不是用来控制循环的，而是直接调到lable处从新执行
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == 2 && j == 2 {
				goto lable
			}
			fmt.Printf("%v,%v\n", i, j)
		}
	}
}
*/