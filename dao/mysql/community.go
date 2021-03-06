package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in database")
			err = nil
		}

	}

	return
}

// GetCommunityDetail 根据id获取社区详情
func GetCommunityDetail(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := "select community_id,community_name,introduction,create_time from community where community_id = ?"
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is invalid id in database")
			err = nil
		}
	}
	return community, err
}

// GetCommunityById 获取社区详情
func GetCommunityById(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := "select community_id,community_name,introduction,create_time from community where community_id = ?"
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			// 打印错误日志
			zap.L().Debug("sql", zap.Any(sqlStr, id))

			zap.L().Warn("there is invalid id in database")
			err = nil
		}
	}
	return community, err
}
