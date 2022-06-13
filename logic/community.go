package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

//GetCommunityList  获取社区列表
func GetCommunityList() ([]*models.Community, error) {
	communityList, err := mysql.GetCommunityList()
	if err != nil {
		zap.L().Error("getCommunityList error", zap.Error(err))
		return nil, err
	}
	return communityList, nil
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	communityDetail, err := mysql.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("getCommunityDetail error", zap.Error(err))
		return nil, err
	}
	return communityDetail, nil
}
