/*package main

import (
	"bytes"
	"fmt"
)

func testcontains() { //判断b是否包含某个字节切片
	s := "hello world"
	b := []byte(s)
	b2 := bytes.Contains(b, []byte("hello"))
	fmt.Printf("b2: %v\n", b2)
}

func testcount() { //测试某个切片在母切片的数量
	s := "hello world"
	b := []byte(s)
	i := bytes.Count(b, []byte("h")) //1
	j := bytes.Count(b, []byte("k")) //0
	n := bytes.Count(b, []byte("l")) //3
	fmt.Println(i, j, n)
}

func testrepeat() { //返回一个重复指定次数的[]byte
	b := []byte("hi")
	fmt.Println(bytes.Repeat(b, 0))         //byte按照ascii数值输出了
	fmt.Printf("%s\n", bytes.Repeat(b, 1))  //一种方法，按字符串输出
	fmt.Println(string(bytes.Repeat(b, 1))) //转换成字符串输出
	fmt.Println(bytes.Repeat(b, 1))
	fmt.Println(bytes.Repeat(b, 3))
}

func testreplace() { //返回一个s里old被news替换了指定次数的[]byte
	s := []byte("hello,world")
	old := []byte("l")
	news := []byte("k")
	b := bytes.Replace(s, old, news, 0) //0不替换，1替换一次，-1能替换多少次就替换多少次
	fmt.Printf("b: %v\n", b)
}

func testrunes() {
	b1 := []byte("你好")               //汉字一个用三个字节
	b2 := bytes.Runes(b1)            //runes后只有一个
	fmt.Println(string(b1), len(b1)) //你好 6
	fmt.Println(string(b2), len(b2)) //你好 2
}

func testjoin() {
	s := [][]byte{[]byte("11"), []byte("22"), []byte("33")} //创建一个元素为字节切片的切片
	sep := []byte(",")
	b := bytes.Join(s, sep) //join第一个参数必须是s类型，第二个是字节切片，返回值是用sep连接起来的字节切片
	fmt.Printf("b: %v\n", string(b))
}

func main() {
	testjoin()
}*/
