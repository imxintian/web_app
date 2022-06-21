package redis

const (
	KeyPrefix             = "web_app:"
	KeyPostTimeZSet       = "post:time"  // 存储帖子的时间戳
	KeyPostScoreZSet      = "post:score" // 存储帖子的积分
	KeyPostVoteZSetPrefix = "post:vote:" // 存储用户及投票类型，参数是post_id
	KeyCommunitySetPF     = "community:" // 存储每个社区下帖子的id
) // end const

// getRedisKey 获取redis key
func getRedisKey(key string) string {
	return KeyPrefix + key
}
