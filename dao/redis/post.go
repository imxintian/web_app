package redis

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
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

// GetPostVoteList  按社区id获取帖子id查询每篇帖子的赞同和反对数

func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 使用 zinterstore 把分区的帖子与帖子分数的zeset 生成一个新的zset
	// 并且按照分数从大到小的顺序指定数量的元素
	// 利用缓冲key减少zinterstore执行的次数
	//社区的key
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityId)))
	// 缓存的key
	key := orderKey + strconv.Itoa(int(p.CommunityId))
	if rdb.Exists(orderKey).Val() < 1 {
		// 不存在分区的key，则创建
		pipeline := rdb.Pipeline()
		rdb.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}

	}
	// 确定查询的起始位置
	start := p.Page * (p.Page - 1)
	end := start + p.PageSize - 1

	//查询  按分数从大到小的的顺序指定数量的元素
	return rdb.ZRevRange(key, start, end).Result()

}
