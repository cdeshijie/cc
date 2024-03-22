/*package main

import "fmt"

func main() {
//空切片和mil切片是不一样的
//空切片var slice = []int{}  var slice = make([]int,0)
//nil切片 var slice []int,nil切片和空切片都可以直接append，
	var s1 = make([]int, 3)
	s2 := make([]float64, 2)
	s3 := []int{1, 2, 3} //切片
	var s4 []int  //定义空切片用var，不要用:=
	//s4 := [...]int{12,213,3123,122}//数组 s4无法使用append函数
	var s5 []int
	var s6 = []int{}

	s1 = append(s1, 12) //添加
	s2 = append(s2, 123)
	s3 = append(s3, 1212)

	s5 = append(s5, 12, 11, 13, 14) //删除第3个元素
	fmt.Printf("s5: %v\n", s5)
	s5 = append(s5[:2], s5[3:]...)
	fmt.Printf("s5: %v\n", s5)

	s6 = append(s6, 1, 2, 3, 4, 5)
}
*/