/* package main

import "os"

func writefile() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0755) //读写也是直接覆盖,加上append就是末尾追加
	f.Write([]byte("hello world"))
	//直接写字符串f.WriteString("hello world")
	f.WriteAt([]byte("hello world"), 5) //从第五个字符之后写入内容
	f.Close()
}

func main() {
	writefile()
}
*/