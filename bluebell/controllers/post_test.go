package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试创建帖子的接口
func TestCreatePostHandler(t *testing.T) {
	//这里是给controller层写单元测试，所以专注于参数校验就行，如果是测试逻辑，就要单独去给logic层写单元测试
	//设置模式为测试
	gin.SetMode(gin.TestMode)
	//创建测试路由
	r := gin.Default()
	//创建url地址
	url := "/api/v1/post"
	//创建post请求
	r.POST(url, CreatePostHandler)
	//创建Post请求中body所携带的json数据
	body := `{
		"community_id":1,
		"title":"test",
		"content":"just a test"
	}`
	//创建一个测试用的访问请求
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	//创建记录测试结果的工具
	w := httptest.NewRecorder()
	//进行测试并将测试结果写入记录工具
	r.ServeHTTP(w, req)
	//判断返回的状态码和结果,在CreatePostHandler中，程序会卡在从c获取用户id这里，因为这里的test只有一个简单的路由，没有上下文获取用户id

	//方法一：所以可以通过返回的错误的状态码和信息来判断是否完成了参数校验这步
	assert.Equal(t, 200, w.Code)

	//方法二：也可以通过判断返回的body中是否含有该有的字符串
	assert.Contains(t, w.Body.String(), "需要登录")

	//方法三：通过查看写好的程序我们可以看到，ResponseError函数会返回一个Response类型的json给前端,这里进行反序列化，解析返给前端的内容看看是否正确
	res := new(Response)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Errorf("json.Unmarshal w.Body failed,err:%v", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}
