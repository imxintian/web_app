package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	salStr := "select community_id,community_name from community"
	if err := db.Select(&communityList, salStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in database")
			err = nil
		}

	}

	return
}
