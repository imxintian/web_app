package logic

import (
	"go.uber.org/zap"
	"web_app/dao/redis"
	"web_app/models"
)

// VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.PostVoteData) error {
	zap.L().Debug("VoteForPost", zap.Int64("userID", userID), zap.Any("p", p)) // 参数输出到日志
	return redis.VoteForPost(userID, p.PostId, float64(p.Direction))
}
