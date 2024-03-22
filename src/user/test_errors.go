/*package main

import (
	"errors"
	"time"
)
//error有两个来源，一个是只返回字符串的goland自带的error，可以用errors.New("字符串")来生成并返回，还有就是自定义一个error
//type error interface{  error是一个接口,里面只有一个方法Error，返回字符串
//	Error() string
//}
//errors包只有一个new函数,new函数传入字符串，返回一个error类型参数

type myerror struct { //如何定义自己的error，先定义结构体变量，然后实现error接口的Error函数实现多态性，之后就可以把这个myerror当做一个error用了
	when time.Time
	what string
}

func (e myerror) Error() string {
	return e.when.String() + "--" + e.what
}

func check(s string) (string, error) { //判断字符串是否为空
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}

func main() {

}*/
