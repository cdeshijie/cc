/*package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string
	Age    int
	Gender bool //true为男
}
type Class struct {
	Id       string
	Students []Student
}

func structtojson() { //json序列化
	s := Student{"张三", 18, true}
	c := Class{"1(2)班", []Student{s, s, s}}
	bytes, err := json.Marshal(c) //json序列化,bytes为字节流
	if err != nil {
		fmt.Printf(",json序列化失败，err: %v\n", err)
		return
	} else {
		fmt.Println("json序列化成功，字节流为：", string(bytes))
	}
}
func jsontostruct() { //json反序列化
	s := Student{"张三", 18, true}
	c := Class{"1(2)班", []Student{s, s, s}}
	bytes, err := json.Marshal(c) //json序列化,bytes为字节流
	if err != nil {
		fmt.Printf(",json序列化失败，err: %v\n", err)
		return
	} else {
		fmt.Println("json序列化成功，字节流为：", string(bytes))
	}
	var c2 Class
	err2 := json.Unmarshal(bytes, &c2) //可能会反序列化失败
	if err2 != nil {
		fmt.Println("反序列化失败,err:", err2)
		return
	}
	fmt.Printf("反序列化成功,c2: %v\n", c2)
}
func main() {
	jsontostruct()
}*/
