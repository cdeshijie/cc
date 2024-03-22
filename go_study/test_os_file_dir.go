/* package main

import (
	"fmt"
	"os"
)

func CreateFile() { //os.Create创建文件操作会把原来的文件替换,返回值为文件指针和错误类型
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

func CreateDir() {
	err := os.Mkdir("test", os.ModePerm) //返回值只有错误，os.Mkdir创建一个名为“test”的目录，os.ModePerm的位置为权限大小，os.ModePerm是最高权限,与文件不同，目录不能覆盖,否则会报错
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	err2 := os.MkdirAll("test/b/c", os.ModePerm) //创建多级目录
	if err2 != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func Remove() {
	err := os.Remove("a.txt") //返回值只有错误，删除当前文件夹下名为a的文件,也可以删除目录
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	err2 := os.RemoveAll("test") //删除test目录本身及目录其下所有的东西
	if err2 != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func wd() {
	dir, err := os.Getwd() //返回值为目录和错误,获得当前函数的工作目录,不是文件所在目录，其值直到e:\goprojects，而不是到项目里的src啥的
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}

	os.Chdir("d:/")          //返回值只有错误，修改工作目录到d盘下，作用目前还不清楚，后面学到再补
	dir2, err2 := os.Getwd() //输出修改后的工作目录
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("dir2: %v\n", dir2)
	}

	s := os.TempDir() //获得临时目录
	fmt.Printf("s: %v\n", s)
}

func Rename() {
	err := os.Rename("a.text", "a.txt") //返回值只有错误，改变文件名称
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func Readfile() {
	b, err := os.ReadFile("a.txt") //返回值为字节数组（切片）和错误
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("b: %v\n", string(b)) //转为字符串输出
	}
}

func Writefile() {
	err := os.WriteFile("a.txt", []byte("hello world"), os.ModePerm) //返回值只有var，没有则创建,第二个参数意思是字节数组，就是切片,有文件则会覆盖原有文件内容
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {
	//CreateFile()
	//CreateDir()
	//wd()
	Rename()
	Writefile()
}
*/