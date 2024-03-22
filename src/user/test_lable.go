/* package main

import "fmt"

func main() { //lable 和for和break的配合是，在执行到continue lable后，调到lable处，然后直接进行离lable最近的for的下一次循环
	//这里是i，所以在221后，i直接等于了3，如果是break lable，那么就是把lable之后的最近一个for循环直接终止

	for i := 0; i < 4; i++ {
	lable:
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if i == 2 && j == 2 && k == 2 {
					break lable
				}
				fmt.Printf("i: %v,j:%v,k=%v\n", i, j, k)
			}
		}
	}
}
*/