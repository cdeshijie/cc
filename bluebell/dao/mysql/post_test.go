package mysql

import (
	"gin_demo2/models"
	"github.com/jmoiron/sqlx"
	"testing"
)

func init() {
	var err error
	dsn := "root:80175703qwe@tcp(127.0.0.1:3306)/bluebell?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(50)
	return
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed,err:%v", err)
	}
	t.Logf("CreatePost insert record into mysql success!")
}
