package logic

import (
	"gin_demo2/dao/mysql"
	"gin_demo2/models"
)

// GetCommunityList 查找所有的community并返回
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (date *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(id)
}
