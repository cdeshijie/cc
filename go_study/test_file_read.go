/* package main

import (
	"fmt"
	"io"
	"os"
)

// 打开文件
func openclose() {
	f, err := os.Open("b.txt") //f为文件指针，只读打开,找不到不会创建而是报错
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
	f.Close()
	os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE, 755) //第二个为权限，为读写，并且指定了找不到就创建，第三个是给创建的文件的权限，755为最高
} //openfile第二个参数os.O_RDWR|os.O_CREATE|os.O_TRUNC,分别意味着读写，没有就创建，先清空原文件内容再写入

//创建文件

func createfile() {
	f, err := os.Create("c.txt") //返回值为文件指针和错误，openfile的调用，具体看源码
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 读文件
func readops() {
	f, _ := os.Open("a.txt")
	for {
		buf := make([]byte, 5) //此时只能读5个字符，不能读全，可以采用for循环读取文件直到结尾
		n, err := f.Read(buf)  //read的读取顺序是，读5个再读5个，一直到文件尾，遇到文件尾会先输出不到5个的字符然后再循环才会报遇到文件尾的错误，
		//hello world长度11，每次读5个需要循环4次，三次输出，一次报错
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
		if err == io.EOF { //判断是否到了文件尾
			break
		}
		if err != nil {
			fmt.Printf("err: %v\n", err)
			break
		}
	}
	f.Close()
}

func readat() { //无论是移动光标还是readat本质都是对光标的操作，使用的都是对光标操作的函数，具体看源码
	f, _ := os.Open("a.txt")
	buf := make([]byte, 5)
	n2, _ := f.ReadAt(buf, 3) //从第四个字符开始读，写入buf
	fmt.Printf("string(n2): %v\n", string(n2))

	buf1 := make([]byte, 5)
	f.Seek(3, 0) //seek是移动文件里的光标，0从开始，1从结尾
	n, _ := f.Read(buf1)
	fmt.Printf("string(n): %v\n", string(n))
}

func main() {
	readops()
	//openclose()
}
*/