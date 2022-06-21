package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"math"
	"strconv"
	"time"
)

//  基于用户投票的相关算法 https://www.ruanyifeng.com/blog/algorithm
// 投一 票：432分 86400/200-1 = 432
/*本项目使用简化的投票算法，
direction=1 投票类型只有两种：
	1. 之前没有投票，现在投赞成票 差值的绝对值 1 +432
	2. 之前投反对票，现在投赞成票 差值的绝对值 2 +432*2
direction=-1 投票类型只有两种：
	1. 之前没有投票，现在投反对票 差值的绝对值 1 -432
	2. 之前投赞成票，现在投反对票 差值的绝对值 2  -432*2
direction=0 投票类型只有一种：
	1. 之前投赞成票，现在取消投票 差值的绝对值 1 -432
	2. 之前投反对票，现在取消投票 差值的绝对值 1 +432

投票限制
	1. 每个帖子自发表之后一星期内允许用户投票，超过一星期后不允许投票，到期后删除保存的key,并存入mysql
*/

const (
	ONEWEEKINSECOND = 7 * 24 * 60 * 60
	scorePerVote    = 432
) // end const

var (
	ErrVoteTimeExpired = errors.New("post vote expired")
	ErrVoteRepeated    = errors.New("duplicate voting is not allowed")
)

func CreatePostVote(postID int64) error {
	pipeline := rdb.TxPipeline()

	// 1. 帖子的时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 2. 帖子的分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()

	return err
}

func VoteForPost(userID int64, postID string, value float64) error {
	// 1. 判断投票限制
	// 获取帖子的时间戳
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > ONEWEEKINSECOND {
		return ErrVoteTimeExpired
	}

	// 2. 更新帖子的投票分数

	// 先查询用户当前的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVoteZSetPrefix+postID), strconv.Itoa(int(userID))).Val()
	// 如果用户这一次投票与上一次投票相同，则不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipeline := rdb.TxPipeline()
	// 分数累加
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), diff*op*scorePerVote, postID)
	if ErrVoteTimeExpired != nil {
		zap.L().Debug("VoteForPost", zap.Int64("userID", userID), zap.String("postID", postID), zap.Float64("value", value), zap.Error(ErrVoteTimeExpired))
	}

	// 3. 记录用户为该帖子的投票数据
	if value == 0 {
		rdb.ZRem(getRedisKey(KeyPostVoteZSetPrefix+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVoteZSetPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
