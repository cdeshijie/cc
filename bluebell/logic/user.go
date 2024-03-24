package logic

//logic层存储的是业务逻辑,具体实现函数再做调用
import (
	"gin_demo2/dao/mysql"
	"gin_demo2/models"
	"gin_demo2/pkg/jwt"
	"gin_demo2/pkg/snowflake"
)

// SignUp 用户注册
func SignUp(p *models.ParmaSignUp) (err error) {
	//1.根据注册信息，比如用户名判断是否已经存在,和数据库相关的操作在dao层实现
	if err = mysql.CheckUserExist(p.Username); err != nil {
		//数据库查询出错的错误
		return err
	}
	//2.生成用户id
	userID := snowflake.GetID()
	//3.保存进数据库,和数据库相关的操作在dao层实现
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)
}

// Login 用户登录
func Login(p *models.ParmaLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//如果登录成功则user会被改变，得到数据库的userid，因为这里传的是指针
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	//生成token
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
