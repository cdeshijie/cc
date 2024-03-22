/*package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id       int
	username string
	password string
}

var db *sql.DB //创建一个sql的全局连接
func initdb() (err error) {
	dsn := "root:80175703qwe@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True" //连接字符串格式如左，root为用户名，80175703qwe为密码，tcp这里可以省略，后面不重要
	//不会效验账号密码是否正确
	//注意，这里不要用:=,我们是给全局变量db赋值，然后在main函数使用
	db, err = sql.Open("mysql", dsn)
	//这里也可以这么写sql.Open("mysql", "root:80175703qwe@/go_db")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3) //设置最大连接时间
	db.SetMaxOpenConns(10)                 //设置最多连接数量
	db.SetMaxIdleConns(10)                 //设置空余最多连接数量
	//尝试与数据库建立连接（检验dsn是否正确）
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return nil
}
func insert() { //插入
	s := "insert into user_tbl (username,password) value(?,?)" //user_tbl表有三个参数，id，username,password，id为主键自动增加
	r, err := db.Exec(s, "lisi", "li123")                      //插入函数
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.LastInsertId() //LastInsertId()求插入的最后一行的ID,但实际上表里的行数未必等于该值
		fmt.Printf("i: %v\n", i)
	}
}
func queryOneRow() { //查询一行数据
	s := "select* from user_tbl where id=?" //根据id查询数据
	var u User
	err := db.QueryRow(s, 1).Scan(&u.id, &u.username, &u.password) //将查到的数据传给提前建立好的结构体,这里数据顺序最好和数据库表里的数据顺序一致
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("u: %v\n", u)
	}
}
func queryManyRow() { //查询多行数据,与单行查询不同
	s := "select* from user_tbl"
	r, err := db.Query(s)
	var u User
	defer r.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for r.Next() {
			r.Scan(&u.id, &u.username, &u.password)
			fmt.Printf("u: %v\n", u)
		}
	}
}
func update() { //根据id修改user_tbl的数据
	s := "update user_tbl set username=?,password=? where id=?"
	r, err := db.Exec(s, "big kite", "123456", 2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Printf("影响的行数为: %v\n", i)
	}
}
func delete() { //根据id去删除user_tbl表里的数据
	s := "delete from user_tbl where id=?"
	r, err := db.Exec(s, 1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Printf("影响的行数为: %v\n", i)
	}
}
func main() {
	err := initdb()
	if err != nil {
		fmt.Printf("初始化失败,err: %v\n", err)
		return
	} else {
		fmt.Println("初始化成功！", "db为：", db)
	}
	insert()
}*/
