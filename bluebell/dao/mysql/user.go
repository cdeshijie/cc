package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"gin_demo2/models"
)

const secret = "liwenzhou.com"

//把和用户相关的操作封装为函数，等待logic层根据业务需求调用

// InsertUser 插入用户到mysql
func InsertUser(user *models.User) (err error) {
	//密码加密
	user.Password = encryptPassword(user.Password)
	//执行sql语句进行插入
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

// CheckUserExist 用户存在返回true
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(*) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// encryptPassword 对密码进行加密,secret用来进行换位运算

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	//查询没问题，但是没查到
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	//查询数据库出了问题
	if err != nil {
		return
	}
	//查询没问题,判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

// GetUserById 根据id获得用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
