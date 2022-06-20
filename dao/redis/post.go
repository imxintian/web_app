package redis

import (
	"web_app/models"
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis中获取帖子id列表
	// 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 确定查询的起始位置
	start := p.PageSize * (p.Page - 1)
	end := start + p.PageSize - 1

	//查询  按分数从大到小的的顺序指定数量的元素
	return rdb.ZRevRange(key, start, end).Result()

}
