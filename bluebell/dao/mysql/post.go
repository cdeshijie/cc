package mysql

import (
	"errors"
	"gin_demo2/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post (post_id,title,content,author_id,community_id) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		return errors.New("MySQL insert failed")
	}
	return
}

// GetPostById 根据帖子id查询帖子详情
func GetPostById(pid int64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id from post where post_id = ?`
	err = db.Get(data, sqlStr, pid)
	return
}

func GetPostList(page int64, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post order by create_time desc limit ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// GetPostListByIDs 根据给定的id列表查询数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id,?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ",")) //strings.Join(ids, ",") 将id列表用逗号拼起来，详情看FIND_IN_SET用法
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
