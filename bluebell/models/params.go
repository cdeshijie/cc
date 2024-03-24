package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// 这个文件专门用来定义接收参数的结构体
// ParmaSignUp binding标签是gin框架做参数校验用的,required要求不能为空，如果是空，则shouldbindjson函数就会报错,eqfiled判断两个字段值是否相等
type ParmaSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParmaLogin 接收登录页面传来的参数
type ParmaLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	//userid,谁登陆就是谁
	PostID    string `json:"post_id" binding:"required"`              //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票还是反对票，赞成是1，反对是-1,取消投票0
}

// ParamPostList 获取列表参数
type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}

type ParamCommunityPostList struct {
	*ParamPostList
	CommunityID int64 `json:"community_id" from:"community_id"`
}
