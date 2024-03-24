package controllers

import (
	"errors"
	"fmt"
	"gin_demo2/dao/mysql"
	"gin_demo2/logic"
	"gin_demo2/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//1.接收和校验参数
	var p models.ParmaSignUp
	if err := c.ShouldBindJSON(&p); err != nil { //shouldbindjson 只能解析参数以及判断数据类型，数据是否符合业务，比如不能为空，它不能判断
		zap.L().Error("SignUp with invalid params", zap.Error(err)) //将错误写入日志
		//判断err是不是validator的错误类型，不是则
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParam)
			return
		}
		//是的话则可以翻译后再传给前端
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	/*手动判断前端传来的参数是否符合业务,手动判断过于麻烦，换成第三方库来进行参数校验
	if len(p.Password) == 0 || len(p.Username) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		zap.L().Error("SignUp with invalid params") //将错误写入日志
		c.JSON(http.StatusOK, gin.H{
			"msg": "signup参数不合法",
		})
		return
	}*/

	//2.业务处理,转给logic层实现
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应,
	ResponseSuccess(c, "注册成功！")
}

func LoginHandler(c *gin.Context) {
	//1.接收和校验参数
	p := new(models.ParmaLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid params", zap.Error(err)) //将错误写入日志
		//判断err是不是validator的错误类型，不是则
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParam)
			return
		}
		//是的话则可以翻译后再传给前端
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务逻辑处理
	//登录成功后生成token，如果失败，则传来的token是“”空字符串
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("Login failed", zap.String("username", p.Username), zap.Error(err)) //记录哪个用户在一直尝试登录
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	//3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), //如果id值大于了1<<53-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}
