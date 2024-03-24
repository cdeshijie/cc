package mysql

import (
	"database/sql"
	"gin_demo2/models"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db!")
			err = nil
			return
		}
	}
	return
}

// GetCommunityDetailByID 根据id查询社区详情
func GetCommunityDetailByID(id int64) (*models.CommunityDetail, error) {
	community := new(models.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction,create_time from community where community_id=?`
	err := db.Get(community, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db!")
			err = ErrorInvalidID
			return community, err
		}
	}
	return community, err
}
