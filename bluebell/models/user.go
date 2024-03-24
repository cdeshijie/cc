package models

// 这个结构体和数据库相关，所以标签是db,这里之所以只有三个属性是因为，id是主键自增，额mail为空，剩下四个有默认值
type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Token    string
}
