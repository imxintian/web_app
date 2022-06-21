package redis

import (
	"github.com/go-redis/redis"
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

// GetPostVoteList  获取帖子id查询每篇帖子的赞同和反对数
func GetPostVoteList(ids []string) ([]int64, error) {
	data := make([]int64, 0, len(ids))
	// 使用redis的pipeline一次性查询多条命令，减少rtt
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVoteZSetPrefix + id)
		pipeline.ZCount(key, "-1", "-1")
	}
	res, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	for _, v := range res {
		data = append(data, v.(*redis.IntCmd).Val())

	}
	return data, nil
}
