package routes

import (
	"gin_demo2/controllers"
	"gin_demo2/logger"
	"gin_demo2/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	/*userid := snowflake.GetID()
	fmt.Println(userid)*/
	r.Static("/static", "./static")
	r.LoadHTMLFiles("./templates/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//开始注册业务
	v1 := r.Group("/api/v1")
	//注册
	v1.POST("/signup", controllers.SignUpHandler)
	//登录
	v1.POST("/login", controllers.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware()) //应用jwt认证中间件
	{
		v1.GET("/community", controllers.CommunityHandler)                 //获得所有的社区id和社区名
		v1.GET("/community/:id", controllers.CommunityDetailHandler)       //根据社区id去获得社区详情
		v1.GET("/community_post", controllers.GetCommunityPostListHandler) //也是取社区的帖子，不过会按照一定的顺序和大小

		v1.POST("/post", controllers.CreatePostHandler)       //创建帖子
		v1.GET("/post/:id", controllers.GetPostDetailHandler) //根据帖子id查询帖子详情
		v1.GET("/posts", controllers.GetPostListHandler)      //传入size和page，从mysql获得帖子列表
		v1.GET("/posts2", controllers.GetPostListHandler2)    //传入size和page和排序方式，按照排序方式从不同的zset取出一定量的帖子
		v1.POST("/vote", controllers.PostVoteController)      //投票
	}

	//JWTAuthMiddleware() 这里面封装了jwt认证操作，如果JWTAuthMiddleware()顺利执行，则说明有token且正确,然后就会执行后面的函数，输出pong
	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "xxx")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
