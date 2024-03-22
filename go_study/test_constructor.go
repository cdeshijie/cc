/* package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) { //定义一个构造函数
	if name == "" {
		return nil, fmt.Errorf("name不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("age不能小于0")
	}
	return &Person{
		name: name,
		age:  age,
	}, nil
}

func main() {
	per, err := NewPerson("tom", 20)
	if err == nil {
		fmt.Printf("per: %v\n", *per)
	}
}
*/