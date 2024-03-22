/* package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// 关于缓冲区，初始化b := bytes.NewBuffer(make([]byte, 0))，b.available可以查看缓冲区可用的空间大小
func test1() { //readerstring的使用,读字符串，直到遇到传入的字节
	r := strings.NewReader("hello world") //创建一个reader来读取字符串,strings.NewReader没有缓冲区操作，不是很方便
	//f, _ := os.Open("a.txt")
	//defer f.Close()
	r2 := bufio.NewReader(r) //bufio包里对reader进行更细致的使用,这里还可以读文件指针，读取文件里的内容，bufio.NewReader支持缓冲区操作
	s, _ := r2.ReadString('\n')
	fmt.Printf("s: %v\n", s)
}

func test2() { //read的使用，传入切片，读取切片大小的字符
	r1 := strings.NewReader("12131asdadasd12321dsaads") //reader读取字符串
	r2 := bufio.NewReader(r1)
	s := make([]byte, 10)
	for {
		n, err := r2.Read(s) //从缓冲区读取数据存入切片,n为读取的数量
		if err == io.EOF {   //三次读完数据，第四次循环读出eof
			break
		} else {
			fmt.Printf("s: %s\n", s[0:n])
		}
	}
}

func test3() { //readbyte的使用
	r1 := strings.NewReader("abcdefg")
	r2 := bufio.NewReader(r1)
	b, _ := r2.ReadByte() //从r2中读取一个字节,第一次读a，然后走向b
	fmt.Printf("b: %c\n", b)
	b2, _ := r2.ReadByte() //从r2再读一个字节，第二次读b，走向c
	fmt.Printf("b2: %c\n", b2)
	r2.UnreadByte() //回退一个字节，从c退到b
	b3, _ := r2.ReadByte()
	fmt.Printf("b3: %c\n", b3)
}

func test4() { //readslice的使用，读取传入的字节以及之前所遇到的字节，写入一个切片并返回
	s := strings.NewReader("abc,def,ghi")
	r := bufio.NewReader(s)
	slice1, _ := r.ReadSlice(',')
	fmt.Printf("slice1: %s\n", slice1)
	slice2, _ := r.ReadSlice(',')
	fmt.Printf("slice1: %c\n", slice2)
	slice3, _ := r.ReadSlice(',')
	fmt.Printf("slice1: %s\n", slice3)
	slice4, _ := r.ReadSlice(',')
	fmt.Printf("slice1: %c\n", slice4)
}

func test5() { //测试writer的使用
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0777) //file结构体实现了writer和reader，所以可以多态的使用
	//b := bytes.NewBuffer(make([]byte, 0))         //定义一个缓冲区，缓冲区用的都是byte的形式来存储和读取或者调用的
	//w := bufio.NewWriter(b)                      //这样写也是没问题的
	defer f.Close()
	w := bufio.NewWriter(f)   //建立文件和writer的关联
	w.WriteString("hello gg") //这里的writestring是写入内存
	w.Flush()                 //将内存里的东西写入与w关联的writer，这里就是file a.txt,flush会清空缓存区
}

func test6() { //测试writer的reset的使用
	b1 := bytes.NewBuffer(make([]byte, 0)) //buffer结构体同样实现了writer和reader接口
	w := bufio.NewWriter(b1)
	w.WriteString("advdas")
	b2 := bytes.NewBuffer(make([]byte, 0))
	w.Reset(b2) //reset可以清空当前的缓冲区，并改变与writer关联的缓冲区对象
	w.WriteString("cccc")
	w.Flush()
	fmt.Println(b1)
	fmt.Println(b2)
}

func test7() {
	r := strings.NewReader("123") //这两行创建一个带缓冲区的reader
	r1 := bufio.NewReader(r)

	b := bytes.NewBuffer(make([]byte, 0)) //这两行创建带缓冲区的writer
	w := bufio.NewWriter(b)

	rw := bufio.NewReadWriter(r1, w) //创建一个ReadWriter，参数是一个带缓冲区的reader和一个带缓冲区的writer
	s, _ := rw.ReadString('\n')      //从自己的reader里的缓冲区读字符串
	fmt.Printf("s: %v\n", string(s))
	rw.WriteString("456") //往自己的writer里的缓冲区里写字符串
	rw.Flush()
	fmt.Println(b)
}

func test8() {
	s := strings.NewReader("123 45 789 0")
	sc := bufio.NewScanner(s)
	sc.Split(bufio.ScanWords) //以某种方式进行划分，这里是按单词划分就是遇到空格算是一个单词,实际上是对token的操作，token理解为字符串的某种标记,也可以是按行，按单个字符等等
	for sc.Scan() {           //Scan返回bool值，只要token没有结束，即没有到字符串尾
		fmt.Println(sc.Text()) //输出按token划分好的内容
	}
}

func main() {
	test8()
}
*/