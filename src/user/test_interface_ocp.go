/* package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
	age  int
}

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() { //cat实现pet接口
	fmt.Println("cat eat")
}
func (cat Cat) sleep() {
	fmt.Println("cat sleep")
}

func (dog Dog) eat() { //dog实现pet接口
	fmt.Println("dog eat")
}
func (dog Dog) sleep() {
	fmt.Println("dog sleep")
}

type Person struct { //人结构体
	name string
}

func (person Person) Care(pet Pet) { //定义一个人照顾动物的方法，由于cat和dog都实现了pet接口，所以这里cat和dog都可以作为care的参数
	pet.eat()
	pet.sleep()
} //这里pet的使用体现了ocp原则，允许扩展，不允许修改，即再来一个pig结构体也可以先实现eat和sleep，从而可以作为care的参数，而不用去修改care函数

func main() {
	dog := Dog{}
	cat := Cat{}
	per := Person{"tom"} //一行可以省略逗号，如果写成列式，则tom后面要加逗号
	per.Care(dog)
	per.Care(cat)
}
*/